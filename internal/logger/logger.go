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
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
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
		debugLogger: log.New(debugOutput, Blue+"DEBUG: "+Reset, log.Ldate|log.Ltime),
		infoLogger:  log.New(os.Stdout, Green+"INFO: "+Reset, log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, Red+"ERROR: "+Reset, log.Ldate|log.Ltime|log.Lshortfile),
	}

	return logger
}

func (l *Logger) Debug(source, message string) {
	l.debugLogger.Printf("[%s] %s", source, message)
}

func (l *Logger) Info(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", source, fmt.Sprint(v...))
	l.infoLogger.Println(msg)
}

func (l *Logger) Error(source string, v ...interface{}) {
	msg := fmt.Sprintf("[%s] %s", source, fmt.Sprint(v...))
	l.errorLogger.Println(msg)
}

func (l *Logger) Writer() io.Writer {
	return l.infoLogger.Writer()
}
