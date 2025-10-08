#!/bin/bash
# Portainer Branch Selector - Quick Deployment Script
# This script helps you build and deploy Portainer with the branch selector feature

set -e

echo "====================================================="
echo "Portainer Branch Selector - Deployment Helper"
echo "====================================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed. Please install Docker first.${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Docker is installed${NC}"

# Check if docker-compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo -e "${YELLOW}Warning: docker-compose is not installed. Will use 'docker run' instead.${NC}"
    USE_COMPOSE=false
else
    echo -e "${GREEN}✓ docker-compose is installed${NC}"
    USE_COMPOSE=true
fi

echo ""
echo "Select deployment option:"
echo "1. Build from source and deploy"
echo "2. Quick deploy (skip build)"
echo "3. Build Docker image only"
echo "4. Exit"
echo ""

read -p "Enter option (1-4): " OPTION

case $OPTION in
    1)
        echo ""
        echo "Building from source..."
        
        # Check if Node.js is installed
        if ! command -v node &> /dev/null; then
            echo -e "${RED}Error: Node.js is not installed. Please install Node.js 16+ first.${NC}"
            exit 1
        fi
        
        # Check if Go is installed
        if ! command -v go &> /dev/null; then
            echo -e "${RED}Error: Go is not installed. Please install Go 1.25+ first.${NC}"
            exit 1
        fi
        
        # Check if yarn is installed
        if ! command -v yarn &> /dev/null; then
            echo -e "${RED}Error: Yarn is not installed. Please install Yarn first.${NC}"
            exit 1
        fi
        
        echo "Installing dependencies..."
        yarn install
        
        echo "Building frontend..."
        yarn build
        
        echo "Building backend..."
        cd api
        go build -o ../dist/portainer ./cmd/portainer
        cd ..
        
        echo "Copying mustache templates..."
        cp -r mustache-templates dist/
        
        echo -e "${GREEN}✓ Build complete!${NC}"
        
        # Ask if user wants to deploy
        read -p "Deploy now? (y/n): " DEPLOY
        if [ "$DEPLOY" = "y" ] || [ "$DEPLOY" = "Y" ]; then
            OPTION=2
        else
            echo "Build artifacts are in ./dist directory"
            exit 0
        fi
        ;;
    
    2)
        echo ""
        echo "Deploying Portainer..."
        ;;
    
    3)
        echo ""
        echo "Building Docker image..."
        make build-image TAG=branch-selector
        echo -e "${GREEN}✓ Docker image built: portainerci/portainer-ce:branch-selector${NC}"
        exit 0
        ;;
    
    4)
        echo "Exiting..."
        exit 0
        ;;
    
    *)
        echo -e "${RED}Invalid option${NC}"
        exit 1
        ;;
esac

if [ "$OPTION" = "2" ]; then
    # Get port configuration
    echo ""
    read -p "HTTP Port (default: 9000): " HTTP_PORT
    HTTP_PORT=${HTTP_PORT:-9000}
    
    read -p "HTTPS Port (default: 9443): " HTTPS_PORT
    HTTPS_PORT=${HTTPS_PORT:-9443}
    
    read -p "Edge Port (default: 8000): " EDGE_PORT
    EDGE_PORT=${EDGE_PORT:-8000}
    
    # Stop and remove existing container if exists
    if docker ps -a | grep -q portainer; then
        echo "Stopping existing Portainer container..."
        docker stop portainer 2>/dev/null || true
        docker rm portainer 2>/dev/null || true
    fi
    
    if [ "$USE_COMPOSE" = true ]; then
        # Create docker-compose.yml
        cat > docker-compose.yml <<EOF
version: '3.8'

services:
  portainer:
    image: portainerci/portainer-ce:branch-selector
    container_name: portainer
    restart: always
    ports:
      - "${HTTP_PORT}:9000"
      - "${HTTPS_PORT}:9443"
      - "${EDGE_PORT}:8000"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data

volumes:
  portainer_data:
EOF
        
        echo "Starting Portainer with docker-compose..."
        docker-compose up -d
    else
        # Use docker run
        echo "Starting Portainer with docker run..."
        docker run -d \
            -p ${HTTP_PORT}:9000 \
            -p ${HTTPS_PORT}:9443 \
            -p ${EDGE_PORT}:8000 \
            --name=portainer \
            --restart=always \
            -v /var/run/docker.sock:/var/run/docker.sock \
            -v portainer_data:/data \
            portainerci/portainer-ce:branch-selector
    fi
    
    echo ""
    echo -e "${GREEN}✓ Portainer deployed successfully!${NC}"
    echo ""
    echo "Access Portainer at:"
    echo "  HTTP:  http://localhost:${HTTP_PORT}"
    echo "  HTTPS: https://localhost:${HTTPS_PORT}"
    echo ""
    echo "To view logs: docker logs -f portainer"
    echo "To stop: docker stop portainer"
    echo ""
    
    # Wait for Portainer to start
    echo "Waiting for Portainer to start..."
    sleep 5
    
    # Check if container is running
    if docker ps | grep -q portainer; then
        echo -e "${GREEN}✓ Portainer is running${NC}"
        
        # Try to open browser (Linux only)
        if command -v xdg-open &> /dev/null; then
            read -p "Open browser? (y/n): " OPEN_BROWSER
            if [ "$OPEN_BROWSER" = "y" ] || [ "$OPEN_BROWSER" = "Y" ]; then
                xdg-open "http://localhost:${HTTP_PORT}" &
            fi
        fi
    else
        echo -e "${RED}✗ Container failed to start. Check logs with: docker logs portainer${NC}"
        exit 1
    fi
fi

echo ""
echo "====================================================="
echo "Deployment complete!"
echo "====================================================="
echo ""
echo "Next steps:"
echo "1. Navigate to http://localhost:${HTTP_PORT}"
echo "2. Create your admin user"
echo "3. Connect to your Docker environment"
echo "4. Try creating a stack from Git repository!"
echo ""
echo "To test the branch selector feature:"
echo "  - Go to Stacks → Add stack"
echo "  - Select 'Repository' as build method"
echo "  - Enter: https://github.com/portainer/portainer"
echo "  - See the branch dropdown populate automatically!"
echo ""
