//go:build windows

package utils

import (
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTitleW = kernel32.NewProc("SetConsoleTitleW")
)

// SetTitle
// Описание: Функция устанавливает заголовок окна консоли для Windows
// Аргументы: title string - заголовок окна
// Возвращаемые значения: error - ошибка
// example:
// err := utils.SetTitle("AstralBot")
//
//	if err != nil {
//		log.Error("Ошибка установки заголовка окна: %v\n", err)
//	}
func SetTitle(title string) error {
	_, _, err := setConsoleTitleW.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}
