# Branch Selector Feature - Testing Guide

## Feature Overview

This feature adds a dropdown selector for Git branches and tags when creating stacks from Git repositories in Portainer CE. Previously, users had to manually type the reference (e.g., `refs/heads/main`). Now they can select from a dropdown populated with all available branches and tags.

## What Changed

### Before (CE - Community Edition)
- Users had to manually type the full reference path (e.g., `refs/heads/main`)
- No validation or suggestions
- Error-prone and required knowledge of Git reference syntax

### After (This PR)
- Dropdown automatically populated with all branches and tags from the repository
- User-friendly selection instead of manual typing
- Smart ordering: `refs/heads/main` and `refs/heads/master` appear first
- Works exactly like the Business Edition feature

## Testing Steps

### 1. Start Portainer

Using Docker Compose (recommended):
```bash
cd /path/to/portainer
make build-image TAG=branch-selector-test
docker-compose up -d
```

Or using Docker directly:
```bash
docker run -d -p 9000:9000 -p 8000:8000 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  --name portainer \
  portainerci/portainer-ce:branch-selector-test
```

### 2. Access Portainer
- Navigate to `http://localhost:9000` or `http://your-vps-ip:9000`
- Complete the initial setup (create admin user)
- Connect to your Docker environment

### 3. Test the Feature

#### Test Case 1: Public Repository (No Authentication)
1. Navigate to **Stacks** → **Add stack**
2. Name: `test-public-repo`
3. Build method: Select **Repository**
4. Repository URL: `https://github.com/portainer/portainer`
5. **OBSERVE**: After entering a valid URL, the "Repository reference" field should:
   - Show a loading state briefly
   - Transform into a dropdown selector
   - Display all branches and tags from the repository
6. **ACTION**: Click the dropdown
7. **VERIFY**: You should see:
   - `refs/heads/main` (or `refs/heads/master`) at the top
   - All other branches (e.g., `refs/heads/develop`, `refs/heads/feature-xyz`)
   - All tags (e.g., `refs/tags/v2.0.0`, `refs/tags/v2.1.0`)
8. **ACTION**: Select `refs/heads/develop`
9. **ACTION**: Continue with compose file path (e.g., `build/linux/docker-compose.yml`)
10. **ACTION**: Click "Deploy the stack"
11. **VERIFY**: Stack deploys successfully with code from the develop branch

#### Test Case 2: Private Repository (With Authentication)
1. Navigate to **Stacks** → **Add stack**
2. Name: `test-private-repo`
3. Build method: Select **Repository**
4. **ACTION**: Expand "Authentication" section
5. **ACTION**: Select "Username and password"
6. **ACTION**: Enter your Git credentials
7. Repository URL: `https://github.com/your-username/private-repo`
8. **OBSERVE**: Dropdown should populate with branches from private repo
9. **ACTION**: Select a branch
10. **VERIFY**: Stack configuration continues normally

#### Test Case 3: Invalid Repository URL
1. Navigate to **Stacks** → **Add stack**
2. Name: `test-invalid`
3. Build method: Select **Repository**
4. Repository URL: `https://github.com/nonexistent/repo-that-does-not-exist`
5. **OBSERVE**: 
   - URL validation should fail
   - Error message should appear
   - Dropdown should show default option: `refs/heads/main`
6. **VERIFY**: Cannot proceed with invalid URL

#### Test Case 4: Repository with Many Refs
1. Use a repository with many branches/tags (e.g., `https://github.com/kubernetes/kubernetes`)
2. **VERIFY**: 
   - All refs load successfully
   - Dropdown is searchable/filterable
   - Performance is acceptable

#### Test Case 5: TLS Skip Verify
1. Test with a self-hosted Git server with self-signed certificate
2. **ACTION**: Enable "Skip TLS Verification"
3. **VERIFY**: Refs load successfully despite certificate issues

## Expected API Behavior

### API Endpoint: POST /api/gitops/repo/refs

**Request:**
```json
{
  "repository": "https://github.com/portainer/portainer",
  "username": "",
  "password": "",
  "tlsSkipVerify": false
}
```

**Response:**
```json
[
  "refs/heads/main",
  "refs/heads/develop", 
  "refs/heads/feature-branch-1",
  "refs/tags/v2.0.0",
  "refs/tags/v2.1.0"
]
```

**Error Response (Authentication Failed):**
```json
{
  "message": "Invalid git credential",
  "details": "authentication failed"
}
```

## Testing with cURL

You can test the API directly:

```bash
# Get auth token first
TOKEN=$(curl -X POST 'http://localhost:9000/api/auth' \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","password":"your-password"}' | jq -r '.jwt')

# Test the refs endpoint
curl -X POST 'http://localhost:9000/api/gitops/repo/refs' \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "repository": "https://github.com/portainer/portainer",
    "tlsSkipVerify": false
  }' | jq
```

## UI Components to Verify

### Repository Reference Field
- **Component**: `RefSelector` (in `app/react/portainer/gitops/RefField/RefSelector.tsx`)
- **Type**: Dropdown/Select
- **Location**: Stack creation form, under "Repository URL"
- **Behavior**:
  - Disabled until valid repository URL is entered
  - Shows loading spinner while fetching refs
  - Populated with all refs from repository
  - Default selection: first ref in list (typically `refs/heads/main`)

### Visual Appearance
The dropdown should look like a standard Portainer select component:
- Blue/primary color scheme
- Clear label: "Repository reference"
- Helpful tooltip explaining the syntax
- Required field indicator (*)

## Edge Cases to Test

1. **Empty Repository**: Repository with no branches or tags
   - Should show default: `refs/heads/main`

2. **Network Timeout**: Slow network or large repository
   - Should show loading state
   - Should timeout gracefully

3. **Mixed Refs**: Repository with both branches and tags
   - All should be listed
   - Branches and tags clearly identifiable by prefix

4. **Special Characters**: Branch names with special characters
   - Should display correctly
   - Should be selectable

## Performance Expectations

- **Initial Load**: < 2 seconds for typical repository
- **Large Repos**: < 5 seconds for repos with 100+ refs
- **Caching**: Subsequent loads should be instant (backend caching)

## Troubleshooting

### Dropdown Not Appearing
1. Check browser console for errors
2. Verify API endpoint is registered: Check network tab for `/api/gitops/repo/refs`
3. Ensure backend is built with the new handler

### Empty Dropdown
1. Check repository URL is valid and accessible
2. Verify authentication credentials (for private repos)
3. Check backend logs for errors

### Old Behavior (Text Input)
1. Clear browser cache
2. Hard refresh (Ctrl+Shift+R / Cmd+Shift+R)
3. Verify correct build artifacts are deployed

## Success Criteria

✅ Dropdown appears when valid repository URL is entered
✅ All branches and tags are listed
✅ Selection updates the reference field value
✅ Stack deployment works with selected reference
✅ Works with both public and private repositories
✅ Graceful error handling for invalid URLs
✅ No console errors
✅ No security vulnerabilities (CodeQL passed)

## Screenshots Required

1. **Stack creation page** - showing the repository URL field
2. **Repository reference dropdown** - closed state
3. **Repository reference dropdown** - open with list of branches/tags
4. **Selected branch** - showing a selected branch in the dropdown
5. **Complete form** - ready to deploy with git repository
6. **Network tab** - showing the `/api/gitops/repo/refs` call

## Deployment Checklist for VPS

- [ ] Backend binary built successfully
- [ ] Frontend assets compiled
- [ ] Docker image created
- [ ] Image pushed to registry (if applicable)
- [ ] VPS accessible on required ports (9000, 9443, 8000)
- [ ] Docker installed on VPS
- [ ] Data volume configured for persistence
- [ ] SSL/TLS certificates configured (for HTTPS)
- [ ] Firewall rules updated
- [ ] Test with sample repository
- [ ] Backup existing Portainer data (if upgrading)

## Notes

- This feature maintains backward compatibility
- Existing stacks are not affected
- The API respects existing Git authentication settings
- Reference caching improves performance for frequently accessed repositories
