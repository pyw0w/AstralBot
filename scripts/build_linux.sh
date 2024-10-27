#!/bin/bash
echo "Building for Linux..."

# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/AstralBot_linux_x64

# 32-bit
GOOS=linux GOARCH=386 go build -o bin/AstralBot_linux_x86

# ARM
GOOS=linux GOARCH=arm go build -o bin/AstralBot_linux_arm
GOOS=linux GOARCH=arm64 go build -o bin/AstralBot_linux_arm64

echo "Done!"
