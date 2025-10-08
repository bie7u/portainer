package gitops

import (
	"errors"
	"net/http"

	gittypes "github.com/portainer/portainer/api/git/types"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/request"
	"github.com/portainer/portainer/pkg/libhttp/response"
	"github.com/portainer/portainer/pkg/validate"
)

type repositoryRefsPayload struct {
	Repository        string                         `json:"repository" example:"https://github.com/portainer/portainer" validate:"required"`
	Username          string                         `json:"username" example:"myGitUsername"`
	Password          string                         `json:"password" example:"myGitPassword"`
	AuthorizationType gittypes.GitCredentialAuthType `json:"authorizationType"`
	// TLSSkipVerify skips SSL verification when connecting to the Git repository
	TLSSkipVerify bool `json:"tlsSkipVerify" example:"false"`
}

func (payload *repositoryRefsPayload) Validate(r *http.Request) error {
	if len(payload.Repository) == 0 || !validate.IsURL(payload.Repository) {
		return errors.New("invalid repository URL. Must correspond to a valid URL format")
	}

	return nil
}

// @id GitOperationRepoRefs
// @summary List git repository references
// @description List all references (branches and tags) for a git repository
// @description **Access policy**: authenticated
// @tags gitops
// @security ApiKeyAuth
// @security jwt
// @produce json
// @param body body repositoryRefsPayload true "Repository details"
// @success 200 {array} string "Success"
// @failure 400 "Invalid request"
// @failure 500 "Server error"
// @router /gitops/repo/refs [post]
func (handler *Handler) gitOperationRepoRefs(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	var payload repositoryRefsPayload
	err := request.DecodeAndValidateJSONPayload(r, &payload)
	if err != nil {
		return httperror.BadRequest("Invalid request payload", err)
	}

	refs, err := handler.gitService.ListRefs(
		payload.Repository,
		payload.Username,
		payload.Password,
		payload.AuthorizationType,
		false, // hardRefresh
		payload.TLSSkipVerify,
	)
	if err != nil {
		if errors.Is(err, gittypes.ErrAuthenticationFailure) {
			return httperror.BadRequest("Invalid git credential", err)
		}

		return httperror.InternalServerError("Unable to list git repository references", err)
	}

	return response.JSON(w, refs)
}
