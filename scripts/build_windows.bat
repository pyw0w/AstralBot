@echo off
echo Building for Windows...

:: 32-bit
SET GOOS=windows
SET GOARCH=386
go build -o bin/AstralBot_x86.exe

:: 64-bit
SET GOOS=windows
SET GOARCH=amd64
go build -o bin/AstralBot_x64.exe

:: 64-bit arm
SET GOOS=windows
SET GOARCH=arm64
go build -o bin/AstralBot_arm64.exe

echo Done!
