#!/bin/bash
set -e  # Остановить выполнение при ошибке
echo "Building for all platforms..."

# Путь к файлу .env.example
ENV_FILE=".env.example"

# Проверка на наличие файла .env.example
if [ ! -f "$ENV_FILE" ]; then
    echo "Файл .env.example не найден в корне проекта!"
    exit 1
fi

# Создаем директорию для бинарных файлов, если её нет
if [ ! -d "bin" ]; then
    mkdir -p bin
fi

# Функция для создания бинарного файла и упаковки в ZIP
build_and_zip() {
    local os=$1
    local arch=$2
    local output_name=$3
    local dir_name="bin/${output_name}"
    local zip_name="bin/${output_name}.zip"

    # Устанавливаем переменные среды для сборки и создаем бинарный файл
    GOOS=$os GOARCH=$arch go build -o "${dir_name}/${output_name}"

    # Создаем директорию для этой сборки
    mkdir -p "$dir_name"

    # Копируем бинарный файл и .env.example в эту директорию
    cp "${dir_name}/${output_name}" "$dir_name"
    cp "$ENV_FILE" "$dir_name"

    # Создаем архив
    zip -j "$zip_name" "$dir_name"/*

    # Удаляем временную папку для этой архитектуры
    rm -r "$dir_name"
}

# Сборка и упаковка для Windows
build_and_zip "windows" "386" "AstralBot_windows_x86.exe"
build_and_zip "windows" "amd64" "AstralBot_windows_x64.exe"

# Сборка и упаковка для Linux
build_and_zip "linux" "386" "AstralBot_linux_x86"
build_and_zip "linux" "amd64" "AstralBot_linux_x64"
build_and_zip "linux" "arm" "AstralBot_linux_arm"
build_and_zip "linux" "arm64" "AstralBot_linux_arm64"

# Сборка и упаковка для MacOS
build_and_zip "darwin" "amd64" "AstralBot_mac_x64"
build_and_zip "darwin" "arm64" "AstralBot_mac_arm64"

echo "Done!"
