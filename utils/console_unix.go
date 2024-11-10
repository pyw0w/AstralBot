//go:build !windows
// +build !windows

package utils

import (
	"fmt"
)

// SetConsoleTitle устанавливает заголовок окна консоли для Unix-подобных систем
func SetTitle(title string) error {
	// ANSI escape-последовательность для установки заголовка
	fmt.Printf("\033]0;%s\007", title)
	return nil
}
