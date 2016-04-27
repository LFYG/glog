// +build !windows

package glog

import (
	"os"
)

var (
	posixConsoleTextColorRed    []byte = []byte("\033[31m")
	posixConsoleTextColorYellow []byte = []byte("\033[33m")
	posixConsoleTextColorGreen  []byte = []byte("\033[32m")
	posixConsoleTextColorReset  []byte = []byte("\033[0m")
)

func setConsoleTextColorRed() error {
	_, err := os.Stderr.Write(posixConsoleTextColorRed)
	return err
}

func setConsoleTextColorYellow() error {
	_, err := os.Stderr.Write(posixConsoleTextColorYellow)
	return err
}

func setConsoleTextColorGreen() error {
	_, err := os.Stderr.Write(posixConsoleTextColorGreen)
	return err
}

func setConsoleTextColorReset() error {
	_, err := os.Stderr.Write(posixConsoleTextColorReset)
	return err
}
