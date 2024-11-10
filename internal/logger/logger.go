package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
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
		debug: log.New(debugOutput, Blue+"DEBUG:"+Reset, log.Ldate|log.Ltime),
		info:  log.New(os.Stdout, Green+"INFO:"+Reset, log.Ldate|log.Ltime),
		error: log.New(os.Stderr, Red+"ERROR:"+Reset, log.Ldate|log.Ltime),
	}

	return logger
}

func (l *Logger) Debug(source, message string) {
	l.debug.Printf("[%s] %s", source, message)
}

func (l *Logger) Info(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s]%s", source, fmt.Sprint(v...))
	l.info.Println(msg)
}

func (l *Logger) Error(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s]%s", source, fmt.Sprint(v...))
	l.error.Println(msg)
}
