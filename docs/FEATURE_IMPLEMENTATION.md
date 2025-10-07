# Branch Switching Feature for Portainer

## Overview

This repository contains an enhancement to Portainer CE that adds a **Branch Switching** feature for Docker stacks managed from Git repositories. Users can quickly switch between different Git branches directly from the Portainer UI, and the stack will automatically redeploy from the selected branch.

## What's New

### Features Added

1. **Backend API Endpoints**
   - `GET /api/stacks/{id}/git/branches` - Fetch available branches from the Git repository
   - `POST /api/stacks/{id}/switch-branch` - Switch stack to a different branch and redeploy

2. **Frontend UI Enhancement**
   - Branch selector dropdown in the stack details page
   - Real-time branch loading from Git repository
   - Confirmation dialog before switching branches
   - Loading states and error handling
   - Current branch indicator

3. **Complete Documentation**
   - API documentation with examples
   - User guide for the UI feature
   - VPS deployment guide
   - Integration examples in multiple languages

## Changes Summary

### Backend Changes (Go)

#### New Files Created

1. **`api/http/handler/stacks/stack_switch_branch.go`**
   - Implements the `/stacks/{id}/switch-branch` endpoint
   - Validates stack existence and Git configuration
   - Updates stack's Git reference to the new branch
   - Clones repository from the new branch
   - Triggers stack redeployment
   - Updates database with new configuration

2. **`api/http/handler/stacks/stack_git_branches.go`**
   - Implements the `/stacks/{id}/git/branches` endpoint
   - Fetches available branches using GitService
   - Filters refs to return only branches (refs/heads/*)
   - Handles authentication and authorization

#### Modified Files

1. **`api/http/handler/stacks/handler.go`**
   - Added route for `GET /stacks/{id}/git/branches`
   - Added route for `POST /stacks/{id}/switch-branch`
   - Routes are protected with authentication middleware

### Frontend Changes (AngularJS)

#### Modified Files

1. **`app/portainer/rest/stack.js`**
   - Added `switchBranch` resource method
   - Added `getGitBranches` resource method
   - Both methods map to the new API endpoints

2. **`app/portainer/services/api/stackService.js`**
   - Added `switchBranch(id, branch)` service function
   - Added `getGitBranches(id)` service function
   - Functions provide clean API for controllers to use

3. **`app/portainer/components/forms/stack-redeploy-git-form/stack-redeploy-git-form.controller.js`**
   - Added state management for branches list
   - Added `loadBranches()` method to fetch available branches
   - Added `switchBranch(branch)` method to switch and redeploy
   - Added loading states (branchesLoading, branchSwitching)
   - Integrated confirmation dialog before switching
   - Added imports for modal utilities

4. **`app/portainer/components/forms/stack-redeploy-git-form/stack-redeploy-git-form.html`**
   - Added branch selector dropdown UI
   - Shows all available branches
   - Marks current branch
   - Disables selection of current branch
   - Shows loading spinner during operations
   - Displays helpful information text

### Documentation

#### New Files Created

1. **`docs/BRANCH_SWITCHING_API.md`**
   - Complete API documentation
   - Endpoint descriptions and parameters
   - Request/response examples
   - Authentication and authorization details
   - Integration examples in JavaScript and Python
   - Troubleshooting guide

2. **`docs/BRANCH_SWITCHING_USER_GUIDE.md`**
   - Step-by-step user guide
   - Screenshots descriptions (placeholders for actual screenshots)
   - Best practices
   - Common issues and solutions
   - Related features

3. **`docs/VPS_DEPLOYMENT_GUIDE.md`**
   - Complete deployment guide for VPS
   - System requirements and prerequisites
   - Docker installation steps
   - Building from source
   - Creating Docker image
   - Deployment configurations
   - SSL/TLS setup with Let's Encrypt
   - Maintenance and troubleshooting

## How It Works

### User Flow

1. User navigates to a stack created from a Git repository
2. In the "Redeploy from git repository" section, a branch selector appears
3. The selector automatically loads all available branches from the Git repository
4. User selects a different branch from the dropdown
5. User clicks "Switch Branch" button
6. A confirmation dialog appears warning about the redeployment
7. Upon confirmation, Portainer:
   - Updates the stack's Git configuration with the new branch
   - Clones the repository from the new branch
   - Redeploys the stack with the new code
   - Updates the database
8. User receives a success notification
9. The page reloads showing the updated stack

### Technical Flow

```
Frontend                     Backend                      Git Repository
   |                            |                              |
   |-- GET /git/branches ------>|                              |
   |                            |-- ListRefs() -------------->|
   |                            |<----------------------------|
   |<-- branches[] -------------|                              |
   |                            |                              |
   |-- POST /switch-branch ---->|                              |
   |    {branch: "develop"}     |                              |
   |                            |-- Validate Stack ----------->|
   |                            |-- Update GitConfig           |
   |                            |-- CloneRepository() -------->|
   |                            |<----------------------------|
   |                            |-- DeployStack()              |
   |                            |-- UpdateDatabase()           |
   |<-- Updated Stack ----------|                              |
```

## Installation and Usage

### Quick Start

1. **Clone this repository:**
   ```bash
   git clone https://github.com/bie7u/portainer.git
   cd portainer
   git checkout copilot/add-branch-switching-feature
   ```

2. **Build and deploy** (see detailed instructions in `docs/VPS_DEPLOYMENT_GUIDE.md`):
   ```bash
   # Build Docker image
   docker build -t portainer-custom:latest .
   
   # Run Portainer
   docker run -d -p 9000:9000 -p 8000:8000 \
     --name portainer --restart=always \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v portainer_data:/data \
     portainer-custom:latest
   ```

3. **Access Portainer** at `http://localhost:9000`

4. **Create a stack from Git** and try the branch switching feature!

### Using the Feature

1. Create or navigate to a stack created from a Git repository
2. Scroll to the "Redeploy from git repository" section
3. Use the branch selector dropdown to switch branches
4. Click "Switch Branch" and confirm
5. Wait for the stack to redeploy

## API Examples

### Get Available Branches

```bash
curl -X GET "http://localhost:9000/api/stacks/1/git/branches" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

Response:
```json
{
  "branches": ["main", "develop", "feature/new-feature"]
}
```

### Switch Branch

```bash
curl -X POST "http://localhost:9000/api/stacks/1/switch-branch" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

## Security Considerations

- **Authentication Required**: Both endpoints require valid authentication (API key or JWT)
- **Authorization Checks**: User must have permission to manage the stack
- **Password Sanitization**: Git passwords are never returned in API responses
- **Audit Trail**: All branch switches are logged with user and timestamp
- **Validation**: Extensive validation of stack configuration and branch names

## Compatibility

- **Portainer Version**: Based on Portainer CE (develop branch)
- **Docker**: Requires Docker Engine 20.10+
- **Stack Types**: Supports Docker Compose, Docker Swarm, and Kubernetes stacks
- **Git Providers**: Works with GitHub, GitLab, Bitbucket, Azure DevOps, and self-hosted Git

## Testing

### Manual Testing

1. Create a Git repository with multiple branches
2. Deploy a stack from this repository in Portainer
3. Navigate to the stack details page
4. Verify the branch selector appears and loads branches
5. Select a different branch and click "Switch Branch"
6. Verify the confirmation dialog appears
7. Confirm and verify the stack redeploys successfully
8. Check that the current branch updates

### API Testing

```bash
# Set variables
export PORTAINER_URL="http://localhost:9000"
export API_KEY="your-api-key"
export STACK_ID="1"

# Get branches
curl -X GET "${PORTAINER_URL}/api/stacks/${STACK_ID}/git/branches" \
  -H "Authorization: Bearer ${API_KEY}"

# Switch branch
curl -X POST "${PORTAINER_URL}/api/stacks/${STACK_ID}/switch-branch" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

## Limitations

- Only works with stacks created from Git repositories
- Requires valid Git credentials if repository is private
- Branch switch triggers full stack redeployment (may cause brief downtime)
- Cannot switch to tags or specific commits (use advanced Git reference field instead)

## Future Enhancements

Potential improvements that could be added:

- [ ] Show commit history for each branch
- [ ] Compare differences between branches before switching
- [ ] Preview mode to see what would change
- [ ] Rollback feature to previous branch
- [ ] Branch creation from UI
- [ ] Pull request integration
- [ ] Automated testing before deployment
- [ ] Branch protection rules

## Contributing

This is a feature implementation for Portainer CE. To contribute:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Documentation

- [API Documentation](docs/BRANCH_SWITCHING_API.md)
- [User Guide](docs/BRANCH_SWITCHING_USER_GUIDE.md)
- [VPS Deployment Guide](docs/VPS_DEPLOYMENT_GUIDE.md)

## License

This project maintains the same license as Portainer CE (zlib license). See the main Portainer repository for details.

## Acknowledgments

- Built on top of [Portainer CE](https://github.com/portainer/portainer)
- Uses existing Portainer Git integration infrastructure
- Leverages AngularJS components and patterns from the Portainer codebase

## Support

For issues with this feature:
- Create an issue in this repository
- Include detailed error messages and logs
- Provide steps to reproduce the issue

For general Portainer support:
- Visit [Portainer Documentation](https://docs.portainer.io)
- Join the [Portainer Community](https://www.portainer.io/join-our-community)

## Changelog

### Version 1.0.0 (Initial Release)

**Added:**
- Backend API endpoint for fetching Git branches
- Backend API endpoint for switching branches
- Frontend branch selector in stack details page
- Complete documentation suite
- VPS deployment guide

**Changed:**
- Enhanced stack-redeploy-git-form component with branch switching UI
- Extended StackService with new methods

**Security:**
- All endpoints require authentication
- Proper authorization checks for stack management
- Password sanitization in responses
