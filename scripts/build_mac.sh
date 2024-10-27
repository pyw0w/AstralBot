#!/bin/bash
echo "Building for MacOS..."

# Intel
GOOS=darwin GOARCH=amd64 go build -o bin/AstralBot_mac_x64

# Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o bin/AstralBot_mac_arm64

echo "Done!"
