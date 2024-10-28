#!/bin/bash
export LANG=en_US.UTF-8
echo "Запуск бота AstralBot..."

# Сборка проекта
./scripts/build_all.sh
if [ $? -ne 0 ]; then
    echo "Ошибка при сборке проекта. Пожалуйста, проверьте код."
    exit 1
fi

# Убедитесь, что исполняемый файл существует
if [ ! -f "bin/AstralBot_linux_x64" ]; then
    echo "Исполняемый файл не найден. Пожалуйста, соберите проект."
    exit 1
fi

# Запуск бота
./bin/AstralBot_linux_x64
