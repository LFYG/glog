// +build windows

package glog

import (
	"os"
	"syscall"
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
	hStderr                     = os.Stderr.Fd()
)

func setConsoleTextAttribute(attr uint) error {
	_, _, err := procSetConsoleTextAttribute.Call(hStderr, uintptr(attr))
	return err
}

func setConsoleTextColorRed() error {
	return setConsoleTextAttribute(0x04)
}

func setConsoleTextColorYellow() error {
	return setConsoleTextAttribute(0x06)
}

func setConsoleTextColorGreen() error {
	return setConsoleTextAttribute(0x02)
}

func setConsoleTextColorReset() error {
	return setConsoleTextAttribute(0x07)
}
