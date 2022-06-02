// A Slice is a segment of an array. Slices build on arrays and provide more power, flexibility, and convenience.
/* Slice consists of 3 things
- pointer reference to underlying array
- length of the segment of the array it contains
- the capacity, which is the max size up to which the segment can grow
*/
// Declaration: var s []T

package main

import "fmt"

func main() {
	var nil_slice []string
	fmt.Println(nil_slice) // nil

	a := [5]int{20, 15, 5, 30, 25}

	s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))

	// Initialization
	// make([]T, len, cap) []T
	var make_slice = make([]string, 0, 0)
	fmt.Println(make_slice) // []

	// Literal slice
	var literal_slice = []string{"Go", "TypeScript"}
	fmt.Println(literal_slice) // [Go TypeScript]

	var a2 = [4]string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s1 := a2[0:2] // Select from 0 to 2
	s2 := a2[:3]  // Select first 3
	s3 := a2[2:]  // Select last 2

	fmt.Println("Array:", a2)
	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3)

	// We can create slices from other slices
	var slice = []string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
	}

	s4 := slice[1:]
	fmt.Println("Slice 4:", s4)

	// Built-in functions
	// func copy(dst, src []T) int
	s5 := []string{"a", "b", "c", "d"}
	s6 := make([]string, 0)

	e := copy(s6, s5)

	fmt.Println("Src:", s5)
	fmt.Println("Dst:", s6)
	fmt.Println("Elements:", e)

	// append(slice []T, elems ...T) []T
	s7 := []string{"a", "b", "c", "d"}
	s8 := append(s7, "e", "f")

	fmt.Println("s7:", s7)
	fmt.Println("s8:", s8)
}
