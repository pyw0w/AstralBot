//go:build windows
// +build windows

package utils

import (
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTitleW = kernel32.NewProc("SetConsoleTitleW")
)

// SetConsoleTitle устанавливает заголовок окна консоли для Windows
func SetTitle(title string) error {
	_, _, err := setConsoleTitleW.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}
