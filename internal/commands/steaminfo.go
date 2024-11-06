package commands

import (
	"AstralBot/internal/steam"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	error *log.Logger
}

type LoggerAdapter struct {
	logger *Logger
	source string
}

func (l *LoggerAdapter) Write(p []byte) (n int, err error) {
	// Удаляем лишние символы новой строки
	message := string(p)
	message = strings.TrimSuffix(message, "\n")
	l.logger.Debug(l.source, message)
	return len(p), nil
}

func (l *Logger) Debug(source string, v ...interface{}) {
	message := fmt.Sprint(v...)
	l.debug.Printf("[%s] %s", source, message)
}

func (l *Logger) NewAdapter(source string) io.Writer {
	return &LoggerAdapter{
		logger: l,
		source: source,
	}
}

func NewLogger(debugMode bool) *Logger {
	var debugOutput io.Writer
	if debugMode {
		debugOutput = os.Stdout
	} else {
		debugOutput = io.Discard
	}

	logger := &Logger{
		debug: log.New(debugOutput, "DEBUG:", log.Ldate|log.Ltime),
		info:  log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime),
	}

	return logger
}

func RegisterSteamInfoCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "steam",
		Description: "Получить информацию о профиле Steam",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌ Пожалуйста, укажите Steam ID или URL профиля.", nil
			}

			steamID := args[0]

			// Проверка, является ли steamID URL профиля и извлечение ID
			if strings.HasPrefix(steamID, "https://steamcommunity.com/id/") {
				steamID = strings.TrimPrefix(steamID, "https://steamcommunity.com/id/")
				steamID = strings.TrimSuffix(steamID, "/")
			}

			// Проверка, является ли steamID буквенной формой и преобразование в цифровую
			if len(steamID) <= 17 { // Пример длины буквенного SteamID
				var err error
				steamID, err = steam.ConvertToNumericSteamID(steamID)
				if err != nil {
					return "", err
				}
			}

			gameCount, err := steam.GetOwnedGamesCount(steamID)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("Количество игр: %d", gameCount), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
