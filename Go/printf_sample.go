package main

import "fmt"

func myPrintf (format string, a ...interface {}) {
	fmt.Printf (format, a...)
}

func main () {
	myPrintf ("%d\n", 10)
	myPrintf ("%s\n", "Hello World")
	myPrintf ("%d\t%d\n", 2, 3)
	myPrintf ("Hello World\n")
}
