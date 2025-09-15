#!/bin/bash

# Define the Go version to install
GO_VERSION="1.25.1"

# Detect system architecture
ARCH=$(uname -m)
if [[ "$ARCH" == "x86_64" ]]; then
    GO_ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
    GO_ARCH="arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# Set Go tarball URL
GO_TARBALL="go${GO_VERSION}.linux-${GO_ARCH}.tar.gz"
DOWNLOAD_URL="https://go.dev/dl/${GO_TARBALL}"

# Install directory in Downloads folder
INSTALL_DIR="$HOME/Downloads"

# Download Go tarball
echo "Downloading Go $GO_VERSION for $GO_ARCH..."
wget -q --show-progress "$DOWNLOAD_URL" -O "/tmp/$GO_TARBALL"
if [ $? -ne 0 ]; then
    echo "Failed to download Go tarball."
    exit 1
fi

# Extract Go tarball to the target directory
echo "Extracting Go to $INSTALL_DIR..."
tar -C "$HOME/Downloads" -xzf "/tmp/$GO_TARBALL"


# Determine the shell profile file
SHELL_RC=""
if [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
else
    SHELL_RC="$HOME/.profile"
fi

# Update shell profile to add Go variables if they don't exist
echo "Updating $SHELL_RC..."

# Add GOROOT and GOPATH to the shell config
if ! grep -q "export GOROOT=" "$SHELL_RC"; then
    echo "export GOROOT=$INSTALL_DIR" >> "$SHELL_RC"
fi

if ! grep -q "export GOPATH=" "$SHELL_RC"; then
    echo "export GOPATH=$HOME/go-workspace" >> "$SHELL_RC"
fi

# Add Go binaries to PATH if not already added
if ! grep -q "export PATH=.*\$GOROOT/bin" "$SHELL_RC"; then
    echo 'export PATH=$GOROOT/bin:$PATH' >> "$SHELL_RC"
fi

echo "Go $GO_VERSION installed successfully!"
echo "Please run 'source $SHELL_RC' or restart your terminal to apply changes."
