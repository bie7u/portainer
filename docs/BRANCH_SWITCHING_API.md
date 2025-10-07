# Branch Switching Feature - API Documentation

## Overview

This feature allows users to quickly switch between Git branches for Docker stacks directly from the Portainer UI. When a branch is selected, the stack automatically updates and redeploys from the new branch.

## API Endpoints

### 1. Get Available Branches

**Endpoint:** `GET /api/stacks/{id}/git/branches`

**Description:** Retrieves a list of available branches from the Git repository associated with the stack.

**Path Parameters:**
- `id` (integer, required): Stack identifier

**Response:**
```json
{
  "branches": [
    "main",
    "develop",
    "feature/new-feature",
    "hotfix/bug-fix"
  ]
}
```

**Status Codes:**
- `200 OK`: Success
- `400 Bad Request`: Invalid request (stack not from Git)
- `403 Forbidden`: Permission denied
- `404 Not Found`: Stack not found
- `500 Internal Server Error`: Server error

**Example Request:**
```bash
curl -X GET "https://portainer.example.com/api/stacks/1/git/branches" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

---

### 2. Switch Branch

**Endpoint:** `POST /api/stacks/{id}/switch-branch`

**Description:** Switches the stack to a different Git branch and redeploys it.

**Path Parameters:**
- `id` (integer, required): Stack identifier

**Request Body:**
```json
{
  "branch": "develop"
}
```

**Request Parameters:**
- `branch` (string, required): Name of the branch to switch to (e.g., "develop", "main", "feature/new-feature")

**Response:**
Returns the updated stack object with the new branch configuration.

```json
{
  "Id": 1,
  "Name": "my-stack",
  "Type": 2,
  "EndpointId": 1,
  "GitConfig": {
    "URL": "https://github.com/user/repo",
    "ReferenceName": "develop",
    "ConfigFilePath": "docker-compose.yml",
    "ConfigHash": "abc123..."
  },
  "Status": 1,
  "UpdatedBy": "admin",
  "UpdateDate": 1234567890
}
```

**Status Codes:**
- `200 OK`: Branch switched and stack redeployed successfully
- `400 Bad Request`: Invalid request (stack not from Git, invalid branch name)
- `403 Forbidden`: Permission denied
- `404 Not Found`: Stack not found
- `500 Internal Server Error`: Server error (e.g., failed to clone repository, deployment error)

**Example Request:**
```bash
curl -X POST "https://portainer.example.com/api/stacks/1/switch-branch" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

---

## Authentication

All endpoints require authentication using either:
1. **API Key**: Pass in the `Authorization` header as `Bearer YOUR_API_KEY`
2. **JWT Token**: Pass in the `Authorization` header as `Bearer YOUR_JWT_TOKEN`

You can generate an API key from the Portainer UI under User Settings → API keys.

---

## Authorization

The user must have:
- Access to the environment (endpoint) where the stack is deployed
- Permission to manage stacks (either admin or endpoint admin)
- Access to the specific stack's resource control (if applicable)

---

## Error Handling

All endpoints return errors in the following format:

```json
{
  "message": "Error description",
  "details": "Detailed error message"
}
```

Common errors:
- **Stack is not created from git**: The stack was not created from a Git repository
- **Unable to clone git repository**: Failed to access or clone the Git repository
- **Permission denied**: User does not have sufficient permissions
- **Invalid branch name**: The specified branch does not exist in the repository

---

## Implementation Details

### Branch Switching Process

1. **Validation**: Verifies that the stack exists and is created from a Git repository
2. **Authorization**: Checks user permissions
3. **Branch Update**: Updates the stack's Git configuration with the new branch name
4. **Repository Clone**: Clones the repository from the new branch
5. **Deployment**: Redeploys the stack using the new code
6. **Database Update**: Persists the changes to the database
7. **Response**: Returns the updated stack object

### Supported Stack Types

- Docker Compose stacks (Type 2)
- Docker Swarm stacks (Type 1)
- Kubernetes stacks (Type 3)

---

## Security Considerations

1. **Password Sanitization**: Git authentication passwords are never returned in API responses
2. **TLS Verification**: Respects the stack's TLS skip verification setting
3. **Access Control**: Enforces proper authorization checks before allowing branch switches
4. **Audit Trail**: Updates the `UpdatedBy` and `UpdateDate` fields for tracking

---

## Integration Examples

### Using JavaScript/TypeScript

```javascript
// Get available branches
async function getBranches(stackId) {
  const response = await fetch(`/api/stacks/${stackId}/git/branches`, {
    headers: {
      'Authorization': `Bearer ${apiKey}`
    }
  });
  const data = await response.json();
  return data.branches;
}

// Switch branch
async function switchBranch(stackId, branch) {
  const response = await fetch(`/api/stacks/${stackId}/switch-branch`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${apiKey}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ branch })
  });
  return await response.json();
}
```

### Using Python

```python
import requests

def get_branches(stack_id, api_key):
    url = f"https://portainer.example.com/api/stacks/{stack_id}/git/branches"
    headers = {"Authorization": f"Bearer {api_key}"}
    response = requests.get(url, headers=headers)
    return response.json()["branches"]

def switch_branch(stack_id, branch, api_key):
    url = f"https://portainer.example.com/api/stacks/{stack_id}/switch-branch"
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    data = {"branch": branch}
    response = requests.post(url, headers=headers, json=data)
    return response.json()
```

---

## Testing

### Using curl

```bash
# 1. Get your API key from Portainer UI
API_KEY="your-api-key-here"
PORTAINER_URL="https://portainer.example.com"
STACK_ID="1"

# 2. Get available branches
curl -X GET "${PORTAINER_URL}/api/stacks/${STACK_ID}/git/branches" \
  -H "Authorization: Bearer ${API_KEY}"

# 3. Switch to a different branch
curl -X POST "${PORTAINER_URL}/api/stacks/${STACK_ID}/switch-branch" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

### Using Postman

1. Create a new POST request to `{{baseUrl}}/api/stacks/{{stackId}}/switch-branch`
2. Add Authorization header: `Bearer {{apiKey}}`
3. Set Content-Type to `application/json`
4. Add request body: `{"branch": "develop"}`
5. Send the request

---

## Troubleshooting

### Common Issues

**Issue**: "Stack is not created from git"
- **Solution**: This feature only works with stacks created from Git repositories. Verify that the stack has a GitConfig.

**Issue**: "Unable to clone git repository"
- **Solution**: Check that the Git repository URL is accessible, credentials are correct, and the branch exists.

**Issue**: "Permission denied"
- **Solution**: Ensure the user has admin or endpoint admin privileges, and access to the stack's resource control.

**Issue**: Branch not appearing in the list
- **Solution**: The branch might be a tag or a remote ref. Only branches (refs/heads/*) are shown. Try using the advanced configuration to manually specify the reference.
