//go:build windows
// +build windows

package utils

import (
	"syscall"
	"unsafe"
)

// SetConsoleTitle устанавливает заголовок окна консоли для Windows
func SetConsoleTitle(title string) error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleTitle := kernel32.NewProc("SetConsoleTitleW")
	_, _, err := setConsoleTitle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}
