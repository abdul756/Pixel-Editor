#!/bin/bash

# Function to check Docker status
check_docker() {
    # Try standard socket first
    if docker info > /dev/null 2>&1; then
        return 0
    fi

    # If that fails, look for Docker Desktop socket
    if [ -S /var/run/docker.sock ]; then
        export DOCKER_HOST=unix:///var/run/docker.sock
        if docker info > /dev/null 2>&1; then
            echo "Using system Docker socket."
            return 0
        fi
    fi

    # Check for Docker Desktop socket in home directory
    if [ -S "$HOME/.docker/desktop/docker.sock" ]; then
        export DOCKER_HOST=unix://$HOME/.docker/desktop/docker.sock
        if docker info > /dev/null 2>&1; then
            echo "Using Docker Desktop socket."
            return 0
        fi
    fi

    # Docker is not running
    echo "Error: Docker daemon is not running."
    echo "Please start Docker with one of these commands:"
    echo "  • systemctl start docker (for systemd systems)"
    echo "  • sudo service docker start (for init.d systems)"
    echo "  • Open Docker Desktop if you're using that application"
    
    echo -e "\nWould you like to try running the application natively instead? (y/n)"
    read -r choice
    if [[ "$choice" =~ ^[Yy]$ ]]; then
        echo "Running application natively..."
        ./run-native.sh
        exit $?
    fi
    
    exit 1
}

# Check if Docker daemon is running
check_docker

# Allow local Docker containers to access X server
if ! command -v xhost > /dev/null 2>&1; then
    echo "Error: xhost command not found. Please install X11 utilities package."
    echo "  • sudo apt-get install x11-xserver-utils"
    exit 1
fi

xhost +local:docker

# Build and run the Docker container
echo "Building and starting the container..."
docker-compose up --build

# Clean up X server permissions when done
xhost -local:docker 