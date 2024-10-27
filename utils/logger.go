package utils

import (
	"fmt"
	"io"
	"log"
	"os"
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
	l.logger.Debug(l.source, string(p))
	return len(p), nil
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
		// Убираем лишние отступы в префиксах
		debug: log.New(debugOutput, "DEBUG:", log.Ldate|log.Ltime),
		info:  log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime),
	}

	return logger
}

func (l *Logger) Debug(source string, v ...interface{}) {
	// Форматируем сообщение без лишних пробелов
	msg := fmt.Sprintf("[%s]%s", source, fmt.Sprint(v...))
	l.debug.Println(msg)
}

func (l *Logger) Info(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s]%s", source, fmt.Sprint(v...))
	l.info.Println(msg)
}

func (l *Logger) Error(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s]%s", source, fmt.Sprint(v...))
	l.error.Println(msg)
}
