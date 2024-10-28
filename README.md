# AstralBot

## Описание

AstralBot - это бот, который поддерживает Telegram и Discord.

## Установка

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/yourusername/AstralBot.git
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

## Сборка

### Для Windows

   Запустите:
   ```bash
   .\scripts\build_windows.bat
   ```

### Для Linux/MacOS

   Запустите:
   ```bash
   ./scripts/build_all.sh
   ```

## Запуск

### Для Linux/MacOS

   Запустите:

   ```bash
   ./run.bat
   ```

### Для Windows

   Запустите:

   ```bash
   ./run.bat
   ```


### Для Linux/MacOS

   Запустите:

   ```bash
   ./run.sh
   ```

## Команды

- `!ping` - Проверка работоспособности бота.
- `!help` - Показать доступные команды.
- `!fetch` - Получить данные из внешнего API.

## Примечания

- Убедитесь, что у вас есть доступ к интернету для работы с внешними API.
- Если у вас возникли проблемы, проверьте логи для получения дополнительной информации.