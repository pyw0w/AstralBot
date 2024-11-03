//go:build !windows
// +build !windows

package utils

import (
	"log"
)

// SetConsoleTitle устанавливает заголовок окна консоли для Unix-подобных систем
func SetConsoleTitle(title string) error {
	// ANSI escape-последовательность для установки заголовка
	log.Printf("\033]0;%s\007", title)
	return nil
}
