@echo off
chcp 65001 >nul
echo Запуск бота AstralBot...

:: Сборка проекта
echo Сборка проекта AstralBot...
SET GOOS=windows
SET GOARCH=amd64
go build -o bin\AstralBot_windows_x64.exe
IF ERRORLEVEL 1 (
    echo Ошибка при сборке проекта. Пожалуйста, проверьте код.
    exit /b 1
)

:: Убедитесь, что исполняемый файл существует
echo Проверка файла для запуска AstralBot...
IF NOT EXIST "bin\AstralBot_windows_x64.exe" (
    echo Исполняемый файл не найден. Пожалуйста, соберите проект.
    exit /b 1
)

:: Запуск бота
start "" "bin\AstralBot_windows_x64.exe"
