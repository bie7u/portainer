# Branch Selector Dropdown Feature - Deployment Guide

## Overview
This PR adds a branch/tag selector dropdown to Portainer CE when creating stacks from Git repositories. Previously, this feature was only available in the Business Edition (BE).

## Changes Made

### Backend (Go)
1. **New API Endpoint**: Created `/gitops/repo/refs` endpoint to list all branches and tags from a Git repository
   - File: `api/http/handler/gitops/git_repo_refs.go`
   - Endpoint accepts repository URL, credentials, and TLS settings
   - Returns a list of all refs (branches and tags)

2. **Handler Registration**: Added the new endpoint to the gitops handler
   - File: `api/http/handler/gitops/handler.go`

### Frontend (TypeScript/React)
1. **Enable RefSelector for CE**: Removed the Business Edition check that was restricting the dropdown to BE only
   - File: `app/react/portainer/gitops/RefField/RefField.tsx`
   - Changed from conditional rendering (BE vs CE) to always showing the RefSelector
   - Updated validation to always require repository reference

## How the Feature Works

1. When a user enters a valid Git repository URL, the frontend calls `/gitops/repo/refs`
2. The backend uses the existing Git service to list all branches and tags
3. The frontend displays them in a dropdown, with `refs/heads/main` and `refs/heads/master` prioritized
4. The user can select a branch/tag from the dropdown instead of manually typing it

## Building the Project

### Prerequisites
- Go 1.25.0 or later
- Node.js 16 or later
- Yarn package manager
- Docker (for containerized deployment)

### Build Steps

1. **Install Dependencies**
   ```bash
   # Frontend dependencies
   yarn install
   
   # Backend dependencies are auto-downloaded by Go
   ```

2. **Build Frontend**
   ```bash
   yarn build
   ```

3. **Build Backend**
   ```bash
   cd api
   go build -o portainer ./cmd/portainer
   ```

4. **Build Docker Image** (recommended)
   ```bash
   make build-image TAG=mybranch
   ```

## Deployment Options

### Option 1: Docker Compose (Recommended for VPS)

Create a `docker-compose.yml` file:

```yaml
version: '3.8'

services:
  portainer:
    image: portainerci/portainer-ce:mybranch
    container_name: portainer
    restart: always
    ports:
      - "9000:9000"
      - "9443:9443"
      - "8000:8000"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data

volumes:
  portainer_data:
```

Deploy on your VPS:
```bash
# Build the image first
make build-image TAG=mybranch

# Start Portainer
docker-compose up -d

# Check logs
docker-compose logs -f portainer
```

Access Portainer at: `http://your-vps-ip:9000`

### Option 2: Docker Run

```bash
docker run -d \
  -p 8000:8000 \
  -p 9000:9000 \
  -p 9443:9443 \
  --name=portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  portainerci/portainer-ce:mybranch
```

### Option 3: Standalone Binary

```bash
# Build the backend and frontend
make all

# Copy built files to your VPS
scp -r dist/ user@your-vps:/opt/portainer/

# Run on VPS
cd /opt/portainer
./portainer --data /data
```

## Testing the Feature

1. **Access Portainer**: Navigate to `http://your-vps-ip:9000`
2. **Initial Setup**: Create your admin user
3. **Add Environment**: Connect to your Docker environment
4. **Create Stack from Git**:
   - Go to "Stacks" → "Add stack"
   - Select "Repository" as the build method
   - Enter a Git repository URL (e.g., `https://github.com/portainer/portainer`)
   - Wait for the repository to be validated
   - **NEW**: You'll now see a dropdown with all branches and tags!
   - Select a branch from the dropdown
   - Complete the stack configuration and deploy

## Verification

To verify the feature is working:

1. Check that the API endpoint responds:
   ```bash
   curl -X POST http://localhost:9000/api/gitops/repo/refs \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -d '{
       "repository": "https://github.com/portainer/portainer",
       "tlsSkipVerify": false
     }'
   ```

2. Check the browser console for API calls when entering a repository URL
3. Verify the dropdown appears and is populated with branches

## Troubleshooting

### Branch dropdown not appearing
- Ensure the repository URL is valid and accessible
- Check browser console for any API errors
- Verify the backend is running and accessible

### API returns authentication error
- For private repositories, ensure you provide valid credentials
- Check that the Git credentials are correctly saved

### Build errors
- Clear node_modules and rebuild: `rm -rf node_modules && yarn install`
- Clear Go cache: `go clean -cache`

## Security Considerations

- The feature respects existing authentication mechanisms
- Private repository credentials are handled securely
- TLS verification can be disabled for self-signed certificates (not recommended for production)

## Performance

- Branch/tag lists are cached by the backend
- The cache is invalidated when credentials change or on hard refresh
- Large repositories with many refs may take a few seconds to load

## Support

For issues or questions, please open an issue on the GitHub repository.
