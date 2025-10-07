# Quick Start Guide - Deploy Portainer with Docker Compose

This is a condensed guide to get Portainer running on your VPS in just a few minutes. For detailed documentation, see [DEPLOYMENT.md](./DEPLOYMENT.md).

## Prerequisites

- Docker and Docker Compose installed on your VPS
- Ports 9443 and 9000 available

## Step 1: Install Docker (if not already installed)

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
```

Log out and back in for the group change to take effect.

## Step 2: Deploy Portainer

```bash
# Clone the repository
git clone https://github.com/bie7u/portainer.git
cd portainer

# Start Portainer
docker compose up -d
```

Or, without cloning:

```bash
# Download the docker-compose file
mkdir portainer && cd portainer
wget https://raw.githubusercontent.com/bie7u/portainer/develop/docker-compose.yml

# Start Portainer
docker compose up -d
```

## Step 3: Access Portainer

Open your browser and go to:

```
https://your-vps-ip:9443
```

Or for HTTP:

```
http://your-vps-ip:9000
```

**Important:** Create your admin account within 5 minutes!

## Step 4: Configure Firewall (Optional but Recommended)

```bash
# Allow HTTPS access to Portainer
sudo ufw allow 9443/tcp
# Or, if you prefer HTTP:
sudo ufw allow 9000/tcp
```

## Common Commands

```bash
# View logs
docker compose logs -f

# Stop Portainer
docker compose down

# Update Portainer
docker compose pull && docker compose up -d

# Restart Portainer
docker compose restart
```

## Using SSL/TLS (Recommended for Production)

For a more secure setup with HTTPS only:

```bash
# Generate self-signed certificate
mkdir certs
openssl req -newkey rsa:4096 -nodes -sha256 \
  -keyout certs/portainer.key -x509 -days 365 \
  -out certs/portainer.crt

# Use SSL compose file
docker compose -f docker-compose-ssl.yml up -d
```

## Need Help?

- Full documentation: [DEPLOYMENT.md](./DEPLOYMENT.md)
- Official docs: https://docs.portainer.io
- Issues: https://github.com/portainer/portainer/issues

## What's Next?

After deployment:

1. Log in and create your admin account
2. Connect to local Docker environment (auto-detected)
3. Start managing your containers!
4. Explore app templates for quick deployments
5. Set up additional users and access controls

For advanced configuration, backup strategies, and troubleshooting, see the [complete deployment guide](./DEPLOYMENT.md).
