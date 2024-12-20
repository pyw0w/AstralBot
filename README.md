# AstralBot

[![Go Report Card](https://goreportcard.com/badge/github.com/pyw0w/AstralBot)](https://goreportcard.com/report/github.com/pyw0w/AstralBot)
![GitHub license](https://img.shields.io/github/license/pyw0w/AstralBot)

## 📖 Описание

**AstralBot** - это многофункциональный бот, который поддерживает платформы Telegram и Discord. Он предоставляет различные команды для взаимодействия с пользователями, включая команды для проверки работоспособности, получения информации и другие полезные функции.

## 🚀 Функциональность

- **Поддержка Telegram и Discord:** AstralBot может работать одновременно в Telegram и Discord, предоставляя одинаковый набор команд на обеих платформах.
- **Команды:**
  - `ping`: Проверка работоспособности бота.
  - `help`: Показать доступные команды.
  - `test`: Тестовая команда для проверки функциональности.
  - `steam`: Получение информации о пользователе Steam.
- **Логирование:** Поддержка детализированного логирования для отладки и мониторинга.
- **Веб-интерфейс:** Встроенный веб-сервер для отображения документации API и статических файлов.

## 🛠️ Установка

1. **Клонируйте репозиторий:**

    ```bash
    git clone https://github.com/pyw0w/AstralBot.git
    cd AstralBot
    ```

2. **Убедитесь, что у вас установлен Go (версия 1.16 или выше).**

3. **Установите зависимости:**

    ```bash
    go mod tidy
    ```

4. **Создайте файл `.env` на основе `.env.example` и заполните его вашими токенами API.**

    Пример `.env`:

    ```plaintext
    # API Tokens
    TELEGRAM_TOKEN=your_telegram_token
    DISCORD_TOKEN=your_discord_token

    # Debug mode
    DEBUG_MODE=true

    # Detailed API logs (показывать детальные логи API запросов)
    DETAILED_API_LOGS=false

    # Other settings
    COMMAND_PREFIX=!
    ```

5. **Запустите бота:**

    ```bash
    go run main.go
    ```

## 📜 Лицензия

Этот проект лицензирован под лицензией MIT. См. файл [LICENSE](LICENSE) для подробностей.