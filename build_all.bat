@echo off
echo Сборка проекта...

:: Проверка существования директории bin
IF NOT EXIST "bin" (
    mkdir bin
)

:: Windows
SET GOOS=windows
SET GOARCH=386
go build -o bin/AstralBot_windows_x86.exe
SET GOARCH=amd64
go build -o bin/AstralBot_windows_x64.exe

:: Linux
SET GOOS=linux
SET GOARCH=386
go build -o bin/AstralBot_linux_x86
SET GOARCH=amd64
go build -o bin/AstralBot_linux_x64
SET GOARCH=arm
go build -o bin/AstralBot_linux_arm
SET GOARCH=arm64
go build -o bin/AstralBot_linux_arm64

:: MacOS
SET GOOS=darwin
SET GOARCH=amd64
go build -o bin/AstralBot_mac_x64
SET GOARCH=arm64
go build -o bin/AstralBot_mac_arm64

echo Сборка завершена!
