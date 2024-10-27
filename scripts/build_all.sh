#!/bin/bash
echo "Building for all platforms..."

# Создаем директорию для бинарных файлов если её нет
mkdir -p bin

# Windows
GOOS=windows GOARCH=386 go build -o bin/AstralBot_windows_x86.exe
GOOS=windows GOARCH=amd64 go build -o bin/AstralBot_windows_x64.exe

# Linux
GOOS=linux GOARCH=386 go build -o bin/AstralBot_linux_x86
GOOS=linux GOARCH=amd64 go build -o bin/AstralBot_linux_x64
GOOS=linux GOARCH=arm go build -o bin/AstralBot_linux_arm
GOOS=linux GOARCH=arm64 go build -o bin/AstralBot_linux_arm64

# MacOS
GOOS=darwin GOARCH=amd64 go build -o bin/AstralBot_mac_x64
GOOS=darwin GOARCH=arm64 go build -o bin/AstralBot_mac_arm64

echo "Done!"
