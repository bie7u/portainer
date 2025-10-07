# Deploying Modified Portainer with Branch Switching Feature on VPS

## Overview

This guide will walk you through deploying the modified Portainer (with the branch switching feature) on a personal VPS using Docker.

## Prerequisites

### System Requirements

- VPS with at least 2GB RAM and 20GB storage
- Linux-based OS (Ubuntu 20.04/22.04, Debian 11/12, CentOS 7/8, etc.)
- SSH access to your VPS
- Sudo/root privileges

### Software Requirements

- Docker Engine 20.10 or later
- Docker Compose (optional, for multi-container setups)
- Git (for cloning the repository)

## Step 1: Set Up Your VPS

### 1.1 Connect to Your VPS

```bash
ssh user@your-vps-ip
```

### 1.2 Update System Packages

```bash
sudo apt update && sudo apt upgrade -y  # For Ubuntu/Debian
# OR
sudo yum update -y  # For CentOS/RHEL
```

### 1.3 Install Docker

**For Ubuntu/Debian:**

```bash
# Install prerequisites
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

# Add Docker's official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io

# Start and enable Docker
sudo systemctl start docker
sudo systemctl enable docker

# Add your user to docker group (optional, to run docker without sudo)
sudo usermod -aG docker $USER
```

**For CentOS/RHEL:**

```bash
# Install prerequisites
sudo yum install -y yum-utils

# Add Docker repository
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# Install Docker
sudo yum install -y docker-ce docker-ce-cli containerd.io

# Start and enable Docker
sudo systemctl start docker
sudo systemctl enable docker

# Add your user to docker group (optional)
sudo usermod -aG docker $USER
```

### 1.4 Verify Docker Installation

```bash
docker --version
docker run hello-world
```

## Step 2: Build Modified Portainer

### 2.1 Install Build Dependencies

```bash
# Install Go (required for backend)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install Node.js and Yarn (required for frontend)
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt install -y nodejs
sudo npm install -g yarn

# Install build tools
sudo apt install -y build-essential git
```

### 2.2 Clone the Modified Repository

```bash
cd ~
git clone https://github.com/bie7u/portainer.git
cd portainer
git checkout copilot/add-branch-switching-feature
```

### 2.3 Build Portainer

**Option A: Build Using Docker (Recommended)**

```bash
# Build using the official build container
docker run --rm -v $(pwd):/src -w /src portainer/golang-builder:cross-platform /bin/bash -c "yarn install && yarn build && cd api && go build -o /src/dist/portainer"
```

**Option B: Build Locally**

```bash
# Build frontend
cd ~/portainer
yarn install
yarn build

# Build backend
cd ~/portainer/api
go build -o ../dist/portainer

# Make executable
chmod +x ../dist/portainer
```

### 2.4 Create Docker Image

Create a `Dockerfile` in the portainer directory:

```dockerfile
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY dist/portainer /

VOLUME /data

EXPOSE 9000 9443 8000

ENTRYPOINT ["/portainer"]
```

Build the image:

```bash
cd ~/portainer
docker build -t portainer-custom:latest .
```

## Step 3: Deploy Portainer

### 3.1 Create Data Volume

```bash
docker volume create portainer_data
```

### 3.2 Deploy Portainer Container

**Basic Deployment (HTTP on port 9000):**

```bash
docker run -d \
  -p 9000:9000 \
  -p 8000:8000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  portainer-custom:latest
```

**Secure Deployment (HTTPS on port 9443):**

```bash
docker run -d \
  -p 9443:9443 \
  -p 8000:8000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  portainer-custom:latest \
  --sslcert /path/to/cert.pem \
  --sslkey /path/to/key.pem
```

**With Custom Configuration:**

```bash
docker run -d \
  -p 9443:9443 \
  -p 8000:8000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  -v ~/portainer/certs:/certs \
  portainer-custom:latest \
  --sslcert /certs/cert.pem \
  --sslkey /certs/key.pem \
  --admin-password='$2y$05$8oz75U8m5tI/xT4P0NbSHeVQOEB8v8kWC.tW51y2FJdLrW4hFtYTW'
```

## Step 4: Configure Firewall

### 4.1 UFW (Ubuntu/Debian)

```bash
# Allow Portainer ports
sudo ufw allow 9000/tcp
sudo ufw allow 9443/tcp
sudo ufw allow 8000/tcp

# Allow SSH (if not already allowed)
sudo ufw allow 22/tcp

# Enable firewall
sudo ufw enable
```

### 4.2 Firewalld (CentOS/RHEL)

```bash
# Allow Portainer ports
sudo firewall-cmd --permanent --add-port=9000/tcp
sudo firewall-cmd --permanent --add-port=9443/tcp
sudo firewall-cmd --permanent --add-port=8000/tcp

# Reload firewall
sudo firewall-cmd --reload
```

## Step 5: Access Portainer

### 5.1 Initial Setup

1. Open your browser and navigate to:
   - HTTP: `http://your-vps-ip:9000`
   - HTTPS: `https://your-vps-ip:9443`

2. Create an admin account:
   - Username: admin (or your choice)
   - Password: (choose a strong password)

3. Connect to your local Docker environment (already pre-selected)

### 5.2 Create a Test Stack from Git

1. Click **"Stacks"** in the left menu
2. Click **"Add stack"**
3. Select **"Repository"** as the build method
4. Enter:
   - Repository URL: `https://github.com/docker/awesome-compose` (example)
   - Repository reference: `master`
   - Compose path: `nginx-golang-mysql/compose.yaml`
5. Click **"Deploy the stack"**

### 5.3 Test Branch Switching Feature

1. Navigate to your newly created stack
2. Scroll to the **"Redeploy from git repository"** section
3. You should see the branch selector dropdown
4. The dropdown will populate with available branches
5. Select a different branch (e.g., `develop` if available)
6. Click **"Switch Branch"**
7. Confirm the action
8. Wait for the stack to redeploy

## Step 6: Set Up SSL/TLS (Recommended)

### 6.1 Using Let's Encrypt with Certbot

```bash
# Install Certbot
sudo apt install -y certbot  # Ubuntu/Debian
# OR
sudo yum install -y certbot  # CentOS/RHEL

# Obtain certificate (standalone mode)
sudo certbot certonly --standalone -d your-domain.com

# Certificates will be in:
# /etc/letsencrypt/live/your-domain.com/fullchain.pem
# /etc/letsencrypt/live/your-domain.com/privkey.pem
```

### 6.2 Redeploy Portainer with SSL

```bash
# Stop and remove existing container
docker stop portainer
docker rm portainer

# Deploy with SSL
docker run -d \
  -p 9443:9443 \
  -p 8000:8000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  -v /etc/letsencrypt/live/your-domain.com:/certs:ro \
  portainer-custom:latest \
  --sslcert /certs/fullchain.pem \
  --sslkey /certs/privkey.pem
```

## Step 7: Maintenance and Updates

### 7.1 View Logs

```bash
docker logs portainer
docker logs -f portainer  # Follow logs in real-time
```

### 7.2 Update Portainer

```bash
# Pull latest changes
cd ~/portainer
git pull origin copilot/add-branch-switching-feature

# Rebuild
docker build -t portainer-custom:latest .

# Restart container
docker stop portainer
docker rm portainer
# Run the deployment command again from Step 3.2
```

### 7.3 Backup Data

```bash
# Backup Portainer data
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd):/backup \
  alpine \
  tar czf /backup/portainer-backup-$(date +%Y%m%d).tar.gz /data
```

### 7.4 Restore Data

```bash
# Restore Portainer data
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd):/backup \
  alpine \
  tar xzf /backup/portainer-backup-YYYYMMDD.tar.gz -C /
```

## Step 8: Troubleshooting

### Common Issues

**Issue**: Cannot access Portainer web interface
- Check if container is running: `docker ps | grep portainer`
- Check firewall settings
- Verify port forwarding if behind NAT

**Issue**: Build fails
- Ensure all dependencies are installed
- Check Go and Node.js versions
- Verify internet connectivity for downloading packages

**Issue**: Branch switching feature not working
- Check browser console for errors
- Verify API endpoints are accessible
- Check Portainer logs for backend errors

**Issue**: Container exits immediately
- Check logs: `docker logs portainer`
- Verify volume mounts are correct
- Ensure Docker socket is accessible

## Alternative Deployment Methods

### Using Docker Compose

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  portainer:
    image: portainer-custom:latest
    container_name: portainer
    restart: always
    ports:
      - "9000:9000"
      - "9443:9443"
      - "8000:8000"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data
    command: --sslcert /certs/cert.pem --sslkey /certs/key.pem

volumes:
  portainer_data:
```

Deploy:

```bash
docker-compose up -d
```

### Using Pre-built Image (If Available)

If you push your custom image to a registry:

```bash
# Tag and push
docker tag portainer-custom:latest your-registry/portainer-custom:latest
docker push your-registry/portainer-custom:latest

# Deploy on VPS
docker run -d \
  -p 9000:9000 \
  --name portainer \
  --restart=always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  your-registry/portainer-custom:latest
```

## Security Considerations

1. **Use HTTPS**: Always use SSL/TLS in production
2. **Strong Passwords**: Use strong admin passwords
3. **Firewall**: Restrict access to necessary ports only
4. **Updates**: Keep Docker and Portainer updated
5. **Backups**: Regular backups of Portainer data
6. **Network**: Consider using Docker networks for isolation
7. **Secrets**: Use Docker secrets for sensitive data

## Performance Optimization

1. **Resource Limits**: Set CPU and memory limits for the container
2. **Logging**: Configure log rotation
3. **Monitoring**: Set up monitoring and alerts
4. **Caching**: Use Docker BuildKit for faster builds

## Next Steps

1. Configure remote endpoints (if managing multiple Docker hosts)
2. Set up user authentication and RBAC
3. Configure webhooks for automatic deployments
4. Integrate with CI/CD pipelines
5. Monitor stack deployments and logs

## Support

For issues specific to this modified version:
- Check the GitHub repository: https://github.com/bie7u/portainer
- Review the documentation in the `docs/` folder

For general Portainer support:
- Official documentation: https://docs.portainer.io
- Community forum: https://portainer.io/slack
