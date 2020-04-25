package utils

import "github.com/buger/goterm"

//Printf 打印
func Printf(format string, a ...interface{}) (n int, err error) {
	n1, err1 := goterm.Printf(format, a...)
	goterm.Flush()
	return n1, err1
}

//Println 打印
func Println(a ...interface{}) (n int, err error) {
	n1, err1 := goterm.Println(a...)
	goterm.Flush()
	return n1, err1
}

//MoveCursor ...
func MoveCursor(x int, y int) {
	goterm.MoveCursor(1, 1)
}

//Clear ...
func Clear() {
	goterm.Clear()
}
