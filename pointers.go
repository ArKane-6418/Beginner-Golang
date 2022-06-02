// var x *T

package main

import "fmt"

func main() {
	var p *int

	// nil is the zero value for pointers, interfaces, channels, maps, and slices.
	fmt.Println(p) // nil

	a := 10

	var pa *int = &a

	fmt.Printf("address: %p, value: %d\n", pa, *pa)
	*pa = 20
	fmt.Println("after:", a)

	// New function

	new_ptr := new(int)
	*new_ptr = 100

	// Pointer to pointer
	new_ptr_ptr := &new_ptr

	fmt.Printf("address: %p, value: %d\n", new_ptr, *new_ptr)
	fmt.Printf("address: %p, value: %p, further_value: %d\n", new_ptr_ptr, *new_ptr_ptr, **new_ptr_ptr)

	// Go does not support pointer arithmetic
}

// Pointers as function arguments
func myPointerFunction(ptr *int) int {
	return *ptr
}
