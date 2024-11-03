package utils

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"unsafe"
)

func SetConsoleTitle(title string) error {
	switch runtime.GOOS {
	case "windows":
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		setConsoleTitle := kernel32.NewProc("SetConsoleTitleW")
		_, _, err := setConsoleTitle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
		if err != nil && err.Error() != "The operation completed successfully." {
			return err
		}
	case "darwin", "linux":
		// Эмуляция установки заголовка через последовательности ANSI
		log.Printf("\033]0;%s\007", title)
	default:
		return fmt.Errorf("неподдерживаемая ОС: %s", runtime.GOOS)
	}
	return nil
}
