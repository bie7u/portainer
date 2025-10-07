package gitops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gittypes "github.com/portainer/portainer/api/git/types"
	"github.com/portainer/portainer/api/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

// mockGitService is a mock implementation of portainer.GitService for testing
type mockGitService struct {
	ListRefsFunc func(repositoryURL, username, password string, authType gittypes.GitCredentialAuthType, hardRefresh bool, tlsSkipVerify bool) ([]string, error)
}

func (m *mockGitService) CloneRepository(destination, repositoryURL, referenceName, username, password string, authType gittypes.GitCredentialAuthType, tlsSkipVerify bool) error {
	return nil
}

func (m *mockGitService) LatestCommitID(repositoryURL, referenceName, username, password string, authType gittypes.GitCredentialAuthType, tlsSkipVerify bool) (string, error) {
	return "", nil
}

func (m *mockGitService) ListRefs(repositoryURL, username, password string, authType gittypes.GitCredentialAuthType, hardRefresh bool, tlsSkipVerify bool) ([]string, error) {
	if m.ListRefsFunc != nil {
		return m.ListRefsFunc(repositoryURL, username, password, authType, hardRefresh, tlsSkipVerify)
	}
	return nil, nil
}

func (m *mockGitService) ListFiles(repositoryURL, referenceName, username, password string, authType gittypes.GitCredentialAuthType, dirOnly, hardRefresh bool, includedExts []string, tlsSkipVerify bool) ([]string, error) {
	return nil, nil
}

func Test_gitOperationRepoRefs(t *testing.T) {
	handler := NewHandler(
		testhelpers.NewTestRequestBouncer(),
		testhelpers.NewDatastore(),
		&mockGitService{
			ListRefsFunc: func(repositoryURL, username, password string, authType gittypes.GitCredentialAuthType, hardRefresh bool, tlsSkipVerify bool) ([]string, error) {
				return []string{"refs/heads/main", "refs/heads/develop", "refs/tags/v1.0.0"}, nil
			},
		},
		nil, // FileService - not needed for this test
	)

	t.Run("Returns refs successfully with valid payload", func(t *testing.T) {
		payload := repositoryRefsPayload{
			Repository:        "https://github.com/portainer/portainer",
			Username:          "testuser",
			Password:          "testpass",
			AuthorizationType: gittypes.GitCredentialAuthType_Basic,
			TLSSkipVerify:     false,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/gitops/repo/refs", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var refs []string
		err := json.NewDecoder(rec.Body).Decode(&refs)
		assert.NoError(t, err)
		assert.Len(t, refs, 3)
		assert.Contains(t, refs, "refs/heads/main")
		assert.Contains(t, refs, "refs/heads/develop")
		assert.Contains(t, refs, "refs/tags/v1.0.0")
	})

	t.Run("Returns error with invalid URL", func(t *testing.T) {
		payload := repositoryRefsPayload{
			Repository:    "",
			Username:      "testuser",
			Password:      "testpass",
			TLSSkipVerify: false,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/gitops/repo/refs", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Returns error with authentication failure", func(t *testing.T) {
		failHandler := NewHandler(
			testhelpers.NewTestRequestBouncer(),
			testhelpers.NewDatastore(),
			&mockGitService{
				ListRefsFunc: func(repositoryURL, username, password string, authType gittypes.GitCredentialAuthType, hardRefresh bool, tlsSkipVerify bool) ([]string, error) {
					return nil, gittypes.ErrAuthenticationFailure
				},
			},
			nil,
		)

		payload := repositoryRefsPayload{
			Repository:        "https://github.com/portainer/portainer",
			Username:          "baduser",
			Password:          "badpass",
			AuthorizationType: gittypes.GitCredentialAuthType_Basic,
			TLSSkipVerify:     false,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/gitops/repo/refs", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		failHandler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Returns error with incorrect repository URL", func(t *testing.T) {
		failHandler := NewHandler(
			testhelpers.NewTestRequestBouncer(),
			testhelpers.NewDatastore(),
			&mockGitService{
				ListRefsFunc: func(repositoryURL, username, password string, authType gittypes.GitCredentialAuthType, hardRefresh bool, tlsSkipVerify bool) ([]string, error) {
					return nil, gittypes.ErrIncorrectRepositoryURL
				},
			},
			nil,
		)

		payload := repositoryRefsPayload{
			Repository:        "https://github.com/nonexistent/repo",
			Username:          "testuser",
			Password:          "testpass",
			AuthorizationType: gittypes.GitCredentialAuthType_Basic,
			TLSSkipVerify:     false,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/gitops/repo/refs", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		failHandler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
