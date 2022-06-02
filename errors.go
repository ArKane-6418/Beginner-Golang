// No exception handling in Go
// Built-in error type is just an interface type
/*
type error interface {
    Error() string
}
*/

package main

import (
	"errors"
	"fmt"
)

// Wwe know that error is just an interface so anything can be an error as long as it implements the Error() method which returns an error message string
// Custom error

type DivisionError struct {
	Code int
	Msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

// Sentinel errors are expected Errors that can be checked explicitly in other parts of the code
var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {
	result, err := Divide(4, 0)

	if err != nil {
		var divErr DivisionError

		// errors.Is() examines if the error is a particular error object
		// errors.As() checks whether the error has a specific type
		switch {
		case errors.As(err, &divErr):
			fmt.Println(divErr)
			// Do something with the error
		default:
			fmt.Println("no idea!")
		}

		return
	}

	fmt.Println(result)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		// Using errors.New
		// return 0, errors.New("cannot divide by zero")

		// Another way to construct our errors is by using the fmt.Errorf function.
		// return 0, fmt.Errorf("cannot divide %d by zero", a)

		// Using sentinel errors
		// return 0, ErrDivideByZero

		// Custom error
		return 0, DivisionError{
			Code: 2000,
			Msg:  "cannot divide by zero",
		}
	}

	return a / b, nil
}
