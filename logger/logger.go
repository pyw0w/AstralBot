package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func Info(component, message string) {
	InfoLogger.Printf("[%s] %s", component, message)
}

func Error(component, message string) {
	ErrorLogger.Printf("[%s] %s", component, message)
}
