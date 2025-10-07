# Deploying Portainer on a VPS using Docker Compose

This guide will help you deploy Portainer Community Edition on your VPS (Virtual Private Server) using Docker Compose.

## Prerequisites

Before deploying Portainer, ensure your VPS has:

- **Docker Engine** installed (version 20.10.0 or later)
- **Docker Compose** installed (version 1.29.0 or later)
- At least **512MB RAM** (1GB+ recommended)
- **Port 9443** available (for HTTPS access)
- **Port 9000** available (optional, for HTTP access)
- **Port 8000** available (optional, for Edge agent communication)

### Installing Docker and Docker Compose

If you haven't installed Docker yet, follow the official Docker installation guide for your Linux distribution:

```bash
# For Ubuntu/Debian
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Add your user to the docker group (optional, to run docker without sudo)
sudo usermod -aG docker $USER
# Log out and back in for this to take effect

# Verify installation
docker --version
docker compose version
```

## Quick Start

### 1. Clone or Download the Repository

```bash
git clone https://github.com/bie7u/portainer.git
cd portainer
```

Or download just the docker-compose file:

```bash
mkdir portainer-deployment
cd portainer-deployment
wget https://raw.githubusercontent.com/bie7u/portainer/develop/docker-compose.yml
```

### 2. Deploy Portainer

```bash
docker compose up -d
```

### 3. Access Portainer

Open your browser and navigate to:

```
https://your-vps-ip:9443
```

Or if you prefer HTTP:

```
http://your-vps-ip:9000
```

**Important**: On first access, you'll be prompted to create an admin account. Make sure to do this within 5 minutes of starting Portainer, or the container will shut down for security reasons.

## Deployment Options

### Option 1: Basic Deployment (HTTP + HTTPS)

Use the provided `docker-compose.yml` file in the repository root:

```yaml
version: '3.8'

services:
  portainer:
    image: portainer/portainer-ce:latest
    container_name: portainer
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data
    ports:
      - "9000:9000"
      - "9443:9443"
      - "8000:8000"

volumes:
  portainer_data:
```

### Option 2: HTTPS Only (Recommended for Production)

Use `docker-compose-ssl.yml` for a more secure setup:

```yaml
version: '3.8'

services:
  portainer:
    image: portainer/portainer-ce:latest
    container_name: portainer
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data
      - ./certs:/certs:ro
    ports:
      - "9443:9443"
      - "8000:8000"
    command: --sslcert /certs/portainer.crt --sslkey /certs/portainer.key

volumes:
  portainer_data:
```

To use this configuration:

1. Generate SSL certificates (see SSL/TLS section below)
2. Deploy with: `docker compose -f docker-compose-ssl.yml up -d`

### Option 3: Behind a Reverse Proxy (Nginx/Traefik)

If you're using a reverse proxy, you can expose Portainer only on localhost:

```yaml
version: '3.8'

services:
  portainer:
    image: portainer/portainer-ce:latest
    container_name: portainer
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data
    ports:
      - "127.0.0.1:9000:9000"
      - "127.0.0.1:9443:9443"
    networks:
      - proxy

volumes:
  portainer_data:

networks:
  proxy:
    external: true
```

Then configure your reverse proxy to forward requests to `http://portainer:9000` or `https://portainer:9443`.

## SSL/TLS Configuration

### Using Self-Signed Certificates

Generate self-signed certificates:

```bash
mkdir -p certs
openssl req -newkey rsa:4096 -nodes -sha256 -keyout certs/portainer.key -x509 -days 365 -out certs/portainer.crt
```

### Using Let's Encrypt Certificates

If you have a domain name pointed to your VPS:

```bash
# Install certbot
sudo apt-get update
sudo apt-get install certbot

# Generate certificate
sudo certbot certonly --standalone -d your-domain.com

# Copy certificates to the certs directory
mkdir -p certs
sudo cp /etc/letsencrypt/live/your-domain.com/fullchain.pem certs/portainer.crt
sudo cp /etc/letsencrypt/live/your-domain.com/privkey.pem certs/portainer.key
sudo chown $(whoami):$(whoami) certs/*
```

## Environment Variables

You can customize Portainer's behavior using environment variables. Create a `.env` file:

```bash
# Portainer admin password (bcrypt hash)
# Generate with: docker run --rm httpd:2.4-alpine htpasswd -nbB admin "your-password" | cut -d ":" -f 2
PORTAINER_ADMIN_PASSWORD=

# Snapshot interval (default: 5m)
PORTAINER_SNAPSHOT_INTERVAL=5m

# Log level (DEBUG, INFO, WARN, ERROR)
PORTAINER_LOG_LEVEL=INFO

# Edge agent port
PORTAINER_EDGE_PORT=8000
```

Then reference these in your `docker-compose.yml`:

```yaml
services:
  portainer:
    # ... other configuration ...
    environment:
      - PORTAINER_SNAPSHOT_INTERVAL=${PORTAINER_SNAPSHOT_INTERVAL:-5m}
```

## Firewall Configuration

Make sure to open the necessary ports in your VPS firewall:

```bash
# Using UFW (Ubuntu/Debian)
sudo ufw allow 9443/tcp comment 'Portainer HTTPS'
sudo ufw allow 9000/tcp comment 'Portainer HTTP'
sudo ufw allow 8000/tcp comment 'Portainer Edge Agent'

# Using firewalld (CentOS/RHEL)
sudo firewall-cmd --permanent --add-port=9443/tcp
sudo firewall-cmd --permanent --add-port=9000/tcp
sudo firewall-cmd --permanent --add-port=8000/tcp
sudo firewall-cmd --reload
```

## Managing Your Deployment

### Start Portainer

```bash
docker compose up -d
```

### Stop Portainer

```bash
docker compose down
```

### View Logs

```bash
docker compose logs -f portainer
```

### Update Portainer

```bash
# Pull the latest image
docker compose pull

# Restart with the new image
docker compose up -d

# Remove old images
docker image prune -f
```

### Backup Data

Portainer stores all its data in a Docker volume. To backup:

```bash
# Create a backup directory
mkdir -p backups

# Backup the Portainer data volume
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd)/backups:/backup \
  alpine tar czf /backup/portainer-backup-$(date +%Y%m%d-%H%M%S).tar.gz -C /data .
```

### Restore Data

```bash
# Stop Portainer
docker compose down

# Restore from backup
docker run --rm \
  -v portainer_data:/data \
  -v $(pwd)/backups:/backup \
  alpine sh -c "cd /data && tar xzf /backup/portainer-backup-YYYYMMDD-HHMMSS.tar.gz"

# Start Portainer
docker compose up -d
```

## Troubleshooting

### Portainer Won't Start

1. Check if the ports are already in use:
   ```bash
   sudo netstat -tulpn | grep -E ':(9000|9443|8000)'
   ```

2. Check Docker logs:
   ```bash
   docker compose logs portainer
   ```

3. Ensure Docker socket is accessible:
   ```bash
   ls -l /var/run/docker.sock
   ```

### Cannot Access Portainer Web Interface

1. Verify the container is running:
   ```bash
   docker compose ps
   ```

2. Check firewall rules:
   ```bash
   sudo ufw status
   ```

3. Test connectivity:
   ```bash
   curl -k https://localhost:9443
   ```

### "Admin initialization timeout" Error

If you see this error, it means you didn't create an admin account within 5 minutes. To fix:

```bash
# Stop and remove the container
docker compose down

# Remove the data volume (WARNING: This will delete all Portainer data)
docker volume rm portainer_data

# Start again
docker compose up -d
```

## Security Best Practices

1. **Always use HTTPS in production** - Don't expose port 9000 to the internet
2. **Use strong passwords** - Set a complex admin password
3. **Keep Portainer updated** - Regularly update to get security patches
4. **Limit access** - Use firewall rules to restrict access to trusted IPs
5. **Enable authentication** - Never disable authentication features
6. **Regular backups** - Backup your Portainer data regularly
7. **Use RBAC** - Configure role-based access control for team environments

## Additional Resources

- [Official Portainer Documentation](https://docs.portainer.io)
- [Portainer Installation Guide](https://docs.portainer.io/start/install-ce)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Portainer Community Support](https://www.portainer.io/resources/get-help/get-support)

## Next Steps

After deploying Portainer:

1. **Add Environments** - Connect to other Docker hosts or Kubernetes clusters
2. **Configure Users** - Set up additional users and teams
3. **Deploy Stacks** - Start deploying your applications using Docker Compose
4. **Set up Edge Agent** - Manage remote Docker hosts behind firewalls
5. **Explore Templates** - Use app templates for quick deployments

## Support

If you encounter issues:

- GitHub Issues: https://github.com/portainer/portainer/issues
- Community Slack: https://portainer.io/slack
- Documentation: https://docs.portainer.io
