# Branch Switching Feature - Quick Reference

## Quick Links

- [API Documentation](BRANCH_SWITCHING_API.md)
- [User Guide](BRANCH_SWITCHING_USER_GUIDE.md)
- [VPS Deployment Guide](VPS_DEPLOYMENT_GUIDE.md)
- [Feature Implementation Details](FEATURE_IMPLEMENTATION.md)
- [UI Mockups](UI_MOCKUPS.md)

## At a Glance

### What is it?
A feature that allows quick switching between Git branches for Docker stacks directly from the Portainer UI.

### Why use it?
- Test development branches easily
- Quick rollback to stable versions
- Deploy hotfixes rapidly
- Switch between environments (dev, staging, prod) using branches

### Requirements
- Stack created from a Git repository
- Admin or endpoint admin permissions
- Accessible Git repository

## Quick Start

### For Users (UI)

1. Navigate to your stack in Portainer
2. Find the "Switch to different branch" section
3. Select a branch from the dropdown
4. Click "Switch Branch"
5. Confirm the action
6. Wait for redeployment

**Time:** ~10-30 seconds depending on stack size

### For Developers (API)

```bash
# Get available branches
curl -X GET "https://portainer.example.com/api/stacks/{id}/git/branches" \
  -H "Authorization: Bearer YOUR_API_KEY"

# Switch branch
curl -X POST "https://portainer.example.com/api/stacks/{id}/switch-branch" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

## Files Changed

### Backend (Go)
- ✅ `api/http/handler/stacks/stack_switch_branch.go` (NEW)
- ✅ `api/http/handler/stacks/stack_git_branches.go` (NEW)
- ✅ `api/http/handler/stacks/handler.go` (MODIFIED)

### Frontend (AngularJS)
- ✅ `app/portainer/rest/stack.js` (MODIFIED)
- ✅ `app/portainer/services/api/stackService.js` (MODIFIED)
- ✅ `app/portainer/components/forms/stack-redeploy-git-form/stack-redeploy-git-form.controller.js` (MODIFIED)
- ✅ `app/portainer/components/forms/stack-redeploy-git-form/stack-redeploy-git-form.html` (MODIFIED)

### Documentation
- ✅ `docs/BRANCH_SWITCHING_API.md` (NEW)
- ✅ `docs/BRANCH_SWITCHING_USER_GUIDE.md` (NEW)
- ✅ `docs/VPS_DEPLOYMENT_GUIDE.md` (NEW)
- ✅ `docs/FEATURE_IMPLEMENTATION.md` (NEW)
- ✅ `docs/UI_MOCKUPS.md` (NEW)
- ✅ `docs/QUICK_REFERENCE.md` (NEW - this file)

## API Endpoints

| Method | Endpoint | Purpose |
|--------|----------|---------|
| GET | `/api/stacks/{id}/git/branches` | Get available branches |
| POST | `/api/stacks/{id}/switch-branch` | Switch branch and redeploy |

## Common Commands

### Build Portainer

```bash
# Clone repository
git clone https://github.com/bie7u/portainer.git
cd portainer
git checkout copilot/add-branch-switching-feature

# Build frontend
yarn install
yarn build

# Build backend
cd api
go build -o ../dist/portainer

# Build Docker image
cd ..
docker build -t portainer-custom:latest .
```

### Deploy Portainer

```bash
# Create volume
docker volume create portainer_data

# Run container
docker run -d \
  -p 9000:9000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  portainer-custom:latest
```

### Test API

```bash
# Set variables
export PORTAINER_URL="http://localhost:9000"
export API_KEY="your-api-key"
export STACK_ID="1"

# Get branches
curl "${PORTAINER_URL}/api/stacks/${STACK_ID}/git/branches" \
  -H "Authorization: Bearer ${API_KEY}"

# Switch branch
curl -X POST "${PORTAINER_URL}/api/stacks/${STACK_ID}/switch-branch" \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"branch": "develop"}'
```

## Troubleshooting Quick Fixes

| Problem | Quick Fix |
|---------|-----------|
| Branch selector not showing | Stack not created from Git |
| No branches in dropdown | Check Git credentials |
| Switch fails | Verify branch has docker-compose.yml |
| Permission denied | Need admin/endpoint admin role |
| Cannot access after deployment | Check firewall (ports 9000/9443) |

## Support Stack Types

- ✅ Docker Compose (Type 2)
- ✅ Docker Swarm (Type 1)
- ✅ Kubernetes (Type 3)

## Supported Git Providers

- ✅ GitHub
- ✅ GitLab
- ✅ Bitbucket
- ✅ Azure DevOps
- ✅ Self-hosted Git

## Key Features

- ✅ Automatic branch detection
- ✅ Real-time branch loading
- ✅ Confirmation before switching
- ✅ Loading indicators
- ✅ Success/error notifications
- ✅ Current branch indicator
- ✅ Audit trail (UpdatedBy, UpdateDate)
- ✅ Password sanitization
- ✅ Full authorization checks

## Performance Metrics

| Operation | Expected Time |
|-----------|---------------|
| Load branches | < 2 seconds |
| Switch branch (small stack) | 5-10 seconds |
| Switch branch (large stack) | 10-30 seconds |
| Page reload after switch | < 1 second |

## Security Features

- 🔒 Authentication required (API key or JWT)
- 🔒 Authorization checks (admin/endpoint admin)
- 🔒 Resource control validation
- 🔒 Password sanitization in responses
- 🔒 Audit trail tracking
- 🔒 TLS verification support

## Version Compatibility

| Component | Version |
|-----------|---------|
| Portainer Base | CE (develop branch) |
| Docker Engine | 20.10+ |
| Go | 1.21+ |
| Node.js | 18+ |
| Yarn | Latest |

## Development Workflow

1. Clone repository
2. Create feature branch
3. Make changes
4. Test locally
5. Build and deploy
6. Test in production-like environment
7. Submit PR

## Testing Checklist

- [ ] Backend compiles without errors
- [ ] Frontend builds successfully
- [ ] API endpoints return expected data
- [ ] UI displays branch selector
- [ ] Branch switching works
- [ ] Notifications appear correctly
- [ ] Page reloads after switch
- [ ] Current branch updates
- [ ] Error handling works
- [ ] Loading states display

## Deployment Checklist

- [ ] System requirements met
- [ ] Docker installed
- [ ] Build dependencies installed
- [ ] Repository cloned
- [ ] Code built successfully
- [ ] Docker image created
- [ ] Container deployed
- [ ] Firewall configured
- [ ] SSL/TLS configured (production)
- [ ] Admin account created
- [ ] Test stack created
- [ ] Branch switching tested

## Maintenance Tasks

### Daily
- Monitor logs for errors
- Check stack deployments

### Weekly
- Review audit trail
- Check for updates
- Test critical stacks

### Monthly
- Update SSL certificates (if needed)
- Backup Portainer data
- Review security settings
- Update documentation

## Useful Commands

```bash
# View logs
docker logs portainer
docker logs -f portainer  # Follow

# Restart
docker restart portainer

# Update
docker stop portainer
docker rm portainer
# Run deployment command again

# Backup
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd):/backup \
  alpine \
  tar czf /backup/portainer-backup.tar.gz /data

# Restore
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd):/backup \
  alpine \
  tar xzf /backup/portainer-backup.tar.gz -C /
```

## Getting Help

### Documentation
- Read the full user guide
- Check API documentation
- Review deployment guide
- See UI mockups

### Community
- GitHub Issues
- Portainer Community Slack
- Stack Overflow (tag: portainer)

### Debugging
1. Check browser console (F12)
2. Check Portainer logs
3. Check Docker logs
4. Verify network connectivity
5. Test API endpoints manually

## Best Practices

### For Users
- Test in dev environment first
- Review target branch before switching
- Keep Git credentials updated
- Monitor stack health after switching
- Use descriptive branch names

### For Developers
- Follow existing code patterns
- Write comprehensive tests
- Document changes
- Use semantic versioning
- Keep dependencies updated

### For Admins
- Regular backups
- Monitor performance
- Review audit logs
- Update security settings
- Train users on features

## Known Limitations

- Only works with Git-based stacks
- Requires valid Git credentials
- Full redeployment (may cause downtime)
- Cannot switch to tags directly (use advanced config)
- Branch list cached for 5 minutes

## Roadmap (Future Enhancements)

- [ ] Show commit history
- [ ] Compare branches before switching
- [ ] Preview changes
- [ ] Rollback feature
- [ ] Branch creation from UI
- [ ] PR integration
- [ ] Automated testing

## Credits

Built on top of [Portainer CE](https://github.com/portainer/portainer)

## License

Same as Portainer CE (zlib license)

---

**Last Updated:** 2024
**Version:** 1.0.0
**Branch:** copilot/add-branch-switching-feature
