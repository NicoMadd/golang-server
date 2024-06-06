package util

import "fmt"

const (
	LF   = '\n'
	CR   = '\r'
	CRLF = "\r\n"
	SP   = ' '
)

func ClearConsole() {
	fmt.Printf("\033[H\033[2J")
}
