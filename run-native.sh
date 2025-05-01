#!/bin/bash

# Check if Go is installed
if ! command -v go > /dev/null 2>&1; then
    echo "Error: Go is not installed. Please install Go 1.22 or later."
    echo "Visit https://go.dev/doc/install for installation instructions."
    exit 1
fi

# Check if required dependencies are installed
echo "Checking dependencies..."
if command -v apt-get > /dev/null 2>&1; then
    # For Debian/Ubuntu
    echo "This system uses apt package manager."
    echo "You may need these packages: libgl1-mesa-dev xorg-dev"
    echo "sudo apt-get install libgl1-mesa-dev xorg-dev"
elif command -v pacman > /dev/null 2>&1; then
    # For Arch Linux
    echo "This system uses pacman package manager."
    echo "You may need these packages: libgl mesa xorg-server-devel"
    echo "sudo pacman -S libgl mesa xorg-server-devel"
fi

# Build and run
echo "Building and running PixelEditor..."
cd "$(dirname "$0")"
go run src/pixel.go 