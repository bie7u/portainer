package stacks

import (
	"net/http"
	"time"

	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/git"
	gittypes "github.com/portainer/portainer/api/git/types"
	httperrors "github.com/portainer/portainer/api/http/errors"
	"github.com/portainer/portainer/api/http/security"
	"github.com/portainer/portainer/api/stacks/stackutils"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/request"
	"github.com/portainer/portainer/pkg/libhttp/response"

	"github.com/pkg/errors"
)

type stackSwitchBranchPayload struct {
	Branch string `json:"branch" example:"develop"`
}

func (payload *stackSwitchBranchPayload) Validate(r *http.Request) error {
	if payload.Branch == "" {
		return errors.New("branch is required")
	}
	return nil
}

// @id StackSwitchBranch
// @summary Switch stack to a different git branch
// @description Switch a stack to a different git branch and redeploy
// @description **Access policy**: authenticated
// @tags stacks
// @security ApiKeyAuth
// @security jwt
// @accept json
// @produce json
// @param id path int true "Stack identifier"
// @param body body stackSwitchBranchPayload true "Branch name"
// @success 200 {object} portainer.Stack "Success"
// @failure 400 "Invalid request"
// @failure 403 "Permission denied"
// @failure 404 "Not found"
// @failure 500 "Server error"
// @router /stacks/{id}/switch-branch [post]
func (handler *Handler) stackSwitchBranch(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	stackID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid stack identifier route variable", err)
	}

	stack, err := handler.DataStore.Stack().Read(portainer.StackID(stackID))
	if handler.DataStore.IsErrObjectNotFound(err) {
		return httperror.NotFound("Unable to find a stack with the specified identifier inside the database", err)
	} else if err != nil {
		return httperror.InternalServerError("Unable to find a stack with the specified identifier inside the database", err)
	}

	if stack.GitConfig == nil {
		return httperror.BadRequest("Stack is not created from git", err)
	}

	endpoint, err := handler.DataStore.Endpoint().Endpoint(stack.EndpointID)
	if handler.DataStore.IsErrObjectNotFound(err) {
		return httperror.NotFound("Unable to find the environment associated to the stack inside the database", err)
	} else if err != nil {
		return httperror.InternalServerError("Unable to find the environment associated to the stack inside the database", err)
	}

	if err := handler.requestBouncer.AuthorizedEndpointOperation(r, endpoint); err != nil {
		return httperror.Forbidden("Permission denied to access environment", err)
	}

	securityContext, err := security.RetrieveRestrictedRequestContext(r)
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve info from request context", err)
	}

	// Only check resource control when it is a DockerSwarmStack or a DockerComposeStack
	if stack.Type == portainer.DockerSwarmStack || stack.Type == portainer.DockerComposeStack {
		resourceControl, err := handler.DataStore.ResourceControl().ResourceControlByResourceIDAndType(stackutils.ResourceControlID(stack.EndpointID, stack.Name), portainer.StackResourceControl)
		if err != nil {
			return httperror.InternalServerError("Unable to retrieve a resource control associated to the stack", err)
		}

		if access, err := handler.userCanAccessStack(securityContext, endpoint.ID, resourceControl); err != nil {
			return httperror.InternalServerError("Unable to verify user authorizations to validate stack access", err)
		} else if !access {
			return httperror.Forbidden("Access denied to resource", httperrors.ErrResourceAccessDenied)
		}
	}

	if canManage, err := handler.userCanManageStacks(securityContext, endpoint); err != nil {
		return httperror.InternalServerError("Unable to verify user authorizations to validate stack management", err)
	} else if !canManage {
		errMsg := "Stack management is disabled for non-admin users"
		return httperror.Forbidden(errMsg, errors.New(errMsg))
	}

	var payload stackSwitchBranchPayload
	if err := request.DecodeAndValidateJSONPayload(r, &payload); err != nil {
		return httperror.BadRequest("Invalid request payload", err)
	}

	// Update the branch reference
	stack.GitConfig.ReferenceName = payload.Branch

	repositoryUsername := ""
	repositoryPassword := ""
	repositoryAuthType := gittypes.GitCredentialAuthType_Basic
	if stack.GitConfig.Authentication != nil {
		repositoryUsername = stack.GitConfig.Authentication.Username
		repositoryPassword = stack.GitConfig.Authentication.Password
		repositoryAuthType = stack.GitConfig.Authentication.AuthorizationType
	}

	cloneOptions := git.CloneOptions{
		ProjectPath:   stack.ProjectPath,
		URL:           stack.GitConfig.URL,
		ReferenceName: stack.GitConfig.ReferenceName,
		Username:      repositoryUsername,
		Password:      repositoryPassword,
		AuthType:      repositoryAuthType,
		TLSSkipVerify: stack.GitConfig.TLSSkipVerify,
	}

	clean, err := git.CloneWithBackup(handler.GitService, handler.FileService, cloneOptions)
	if err != nil {
		return httperror.InternalServerError("Unable to clone git repository directory", err)
	}

	defer clean()

	if err := handler.deployStack(r, stack, false, endpoint); err != nil {
		return err
	}

	newHash, err := handler.GitService.LatestCommitID(stack.GitConfig.URL, stack.GitConfig.ReferenceName, repositoryUsername, repositoryPassword, repositoryAuthType, stack.GitConfig.TLSSkipVerify)
	if err != nil {
		return httperror.InternalServerError("Unable get latest commit id", errors.WithMessagef(err, "failed to fetch latest commit id of the stack %v", stack.ID))
	}
	stack.GitConfig.ConfigHash = newHash

	user, err := handler.DataStore.User().Read(securityContext.UserID)
	if err != nil {
		return httperror.BadRequest("Cannot find context user", errors.Wrap(err, "failed to fetch the user"))
	}
	stack.UpdatedBy = user.Username
	stack.UpdateDate = time.Now().Unix()
	stack.Status = portainer.StackStatusActive

	if err := handler.DataStore.Stack().Update(stack.ID, stack); err != nil {
		return httperror.InternalServerError("Unable to persist the stack changes inside the database", errors.Wrap(err, "failed to update the stack"))
	}

	if stack.GitConfig != nil && stack.GitConfig.Authentication != nil && stack.GitConfig.Authentication.Password != "" {
		// Sanitize password in the http response to minimise possible security leaks
		stack.GitConfig.Authentication.Password = ""
	}

	return response.JSON(w, stack)
}
