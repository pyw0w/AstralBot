Write-Host "Building for all platforms..."

# Создаем директорию для бинарных файлов если её нет
New-Item -ItemType Directory -Force -Path bin

# Windows
$env:GOOS = "windows"; $env:GOARCH = "386"; go build -o bin/AstralBot_windows_x86.exe
$env:GOOS = "windows"; $env:GOARCH = "amd64"; go build -o bin/AstralBot_windows_x64.exe

# Linux
$env:GOOS = "linux"; $env:GOARCH = "386"; go build -o bin/AstralBot_linux_x86
$env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o bin/AstralBot_linux_x64
$env:GOOS = "linux"; $env:GOARCH = "arm"; go build -o bin/AstralBot_linux_arm
$env:GOOS = "linux"; $env:GOARCH = "arm64"; go build -o bin/AstralBot_linux_arm64

# MacOS
$env:GOOS = "darwin"; $env:GOARCH = "amd64"; go build -o bin/AstralBot_mac_x64
$env:GOOS = "darwin"; $env:GOARCH = "arm64"; go build -o bin/AstralBot_mac_arm64

Write-Host "Done!"
