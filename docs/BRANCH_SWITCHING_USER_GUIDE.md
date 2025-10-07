# Branch Switching Feature - User Guide

## Overview

The Branch Switching feature allows you to quickly switch between different Git branches for your Docker stacks directly from the Portainer UI. This is useful when you want to:

- Test changes from a development branch
- Switch between different versions of your application
- Roll back to a previous branch
- Deploy hotfixes from a specific branch

## Prerequisites

- The stack must be created from a Git repository
- You must have appropriate permissions (admin or endpoint admin)
- The Git repository must be accessible with the configured credentials

## How to Use the Branch Selector

### 1. Navigate to Your Stack

1. Log in to Portainer
2. Select your environment from the home page
3. Navigate to **Stacks** from the left sidebar
4. Click on the stack you want to manage

### 2. Locate the Branch Selector

The branch selector is located in the **"Redeploy from git repository"** section on the stack details page.

You'll see:
- A dropdown menu showing available branches
- The current branch (marked as "current")
- A "Switch Branch" button

### 3. Switch to a Different Branch

1. Click on the dropdown menu to see all available branches
2. Select the branch you want to switch to
3. Click the **"Switch Branch"** button
4. Confirm the action in the confirmation dialog
   - **Warning**: Switching branches will redeploy the stack and override any local changes

5. Wait for the operation to complete
   - You'll see a loading indicator while the branch is being switched
   - The stack will automatically redeploy from the new branch

6. Upon success:
   - You'll see a success notification
   - The page will reload to show the updated stack
   - The current branch indicator will update

### 4. Verify the Switch

After switching branches, you can verify:
- The current branch is displayed in the branch selector
- The stack status shows as "Active"
- The "Updated by" and "Update date" fields reflect the change
- Your application is running the code from the new branch

## Feature Details

### Branch List

- The dropdown shows all branches from your Git repository
- Branches are listed alphabetically
- The current branch is marked with "(current)" and cannot be selected again
- If no branches are available, the dropdown will be empty

### What Happens During Branch Switch

1. **Validation**: Portainer validates your permissions and the stack configuration
2. **Confirmation**: You must confirm the branch switch action
3. **Repository Clone**: Portainer clones the repository from the selected branch
4. **Deployment**: The stack is redeployed using the new branch's code
5. **Update**: The stack configuration is updated with the new branch
6. **Notification**: You receive a success or error notification

### Important Notes

⚠️ **Warning**: Switching branches will:
- Override any manual changes made to stack files
- Redeploy the stack, which may cause brief downtime
- Update the stack to use the code from the new branch

✅ **Best Practices**:
- Review the code in the target branch before switching
- Ensure the target branch contains a valid docker-compose.yml file
- Test branch switches in a development environment first
- Consider using the "Pull and redeploy" feature for the same branch to get updates

## Advanced Configuration

If you need more control over Git settings while switching branches:

1. Expand the **"Advanced configuration"** section
2. You can modify:
   - Git reference (manually specify branch, tag, or commit)
   - Authentication credentials (if needed)
   - TLS verification settings
   - Relative path settings

## Troubleshooting

### Branch Not Appearing in List

**Possible Causes**:
- The branch doesn't exist in the repository
- There's a connectivity issue with the Git repository
- Authentication credentials are incorrect or expired

**Solutions**:
1. Verify the branch exists in your Git repository
2. Check the repository URL is correct
3. Verify Git authentication credentials
4. Try refreshing the page

### Branch Switch Failed

**Possible Causes**:
- No docker-compose.yml file in the target branch
- Invalid compose file syntax
- Network connectivity issues
- Permission problems

**Solutions**:
1. Check that the target branch contains a valid docker-compose.yml file
2. Verify the compose file syntax is correct
3. Check network connectivity to the Git repository
4. Ensure you have permissions to manage the stack

### Permission Denied

**Possible Causes**:
- You don't have admin or endpoint admin privileges
- Stack management is disabled for regular users
- Resource control restricts your access

**Solutions**:
1. Contact your Portainer administrator
2. Request appropriate permissions
3. Verify you're logged in with the correct account

## Examples

### Switching from Main to Develop Branch

1. Open your stack details page
2. In the branch selector, find "develop" in the dropdown
3. Click "Switch Branch"
4. Confirm the action
5. Wait for the deployment to complete

### Rolling Back to a Previous Branch

1. Identify the branch with the stable version
2. Use the branch selector to switch back to that branch
3. Confirm the rollback
4. Verify your application is working correctly

### Testing a Feature Branch

1. Create a feature branch in your Git repository
2. Push your changes to the feature branch
3. In Portainer, refresh the page to load the new branch
4. Switch to the feature branch using the selector
5. Test your changes
6. Switch back to main when done

## Keyboard Shortcuts

Currently, there are no specific keyboard shortcuts for this feature. You can use standard browser navigation:
- `Tab` to navigate between elements
- `Enter` to activate the selected button
- `Arrow keys` to navigate the dropdown

## Related Features

- **Pull and redeploy**: Update the stack from the same branch
- **Git settings**: Configure repository URL, authentication, and other Git options
- **Auto-update**: Set up automatic updates when the Git repository changes
- **Environment variables**: Manage environment variables for your stack

## Feedback and Support

If you encounter any issues or have suggestions for improving this feature:
- Report bugs on the [Portainer GitHub repository](https://github.com/portainer/portainer/issues)
- Join the [Portainer community](https://www.portainer.io/join-our-community)
- Contact support if you're using Portainer Business Edition
