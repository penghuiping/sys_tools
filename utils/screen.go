package utils

import (
	"fmt"
)

//Printf 打印
func Printf(format string, a ...interface{}) (n int, err error) {
	n1, err1 := fmt.Printf(format, a...)
	return n1, err1
}

//Println 打印
func Println(a ...interface{}) (n int, err error) {
	n1, err1 := fmt.Println(a...)
	return n1, err1
}

//Clear ...
func Clear() {
	fmt.Printf("\033[2J")
}

//MoveCursor ... Move cursor to given position
func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

//MoveCursorUp ... Move cursor up relative the current position
func MoveCursorUp(bias int) {
	fmt.Printf("\033[%dA", bias)
}

//MoveCursorDown  Move cursor down relative the current position
func MoveCursorDown(bias int) {
	fmt.Printf("\033[%dB", bias)
}

//MoveCursorForward Move cursor forward relative the current position
func MoveCursorForward(bias int) {
	fmt.Printf("\033[%dC", bias)
}

//MoveCursorBackward Move cursor backward relative the current position
func MoveCursorBackward(bias int) {
	fmt.Printf("\033[%dD", bias)
}
