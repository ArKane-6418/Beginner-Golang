// In cases where the program cannot continue, we use panic
// func panic(interface{})
// panic stops the normal execution of the current goroutine and control is returned to the caller (repeated until the program exits with panic message and stack trace)

// Possible to regain control using the recover() function and defer
// func recover() interface{}

// Two valid use cases for panic
// - an unrecoverable error
// - developer error
package main

import "fmt"

func main() {
	WillPanic()
}

func handlePanic() {
	data := recover()
	fmt.Println("Recovered:", data)
}

func WillPanic() {
	defer handlePanic()

	panic("Woah")
}
