// var a [n]T
/* Properties
- array's length is part of its type and we cannot resize the array as a result
- arrays are value types
*/

package main

import "fmt"

func main() {

	// Declaration
	var arr [4]int
	fmt.Println(arr) // [0 0 0 0]

	// Initialization
	var arr2 = [4]int{1, 2, 3, 4}
	fmt.Println(arr2)
	// Shorthand initialization
	arr3 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr3)

	// Iteration
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %d, Element: %d\n", i, arr[i])
	}

	// Range iteration
	for i, e := range arr2 {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	// Multi-dimensional
	multiarr := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for i, e := range multiarr {
		fmt.Printf("Index: %d, Element: %d\n", i, e)
	}

	// Arrays are value types
	var a = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var b = a // Copy of a is assigned to b

	b[0] = "Monday"

	fmt.Println(a) // Output: [Mon Tue Wed Thu Fri Sat Sun]
	fmt.Println(b) // Output: [Monday Tue Wed Thu Fri Sat Sun]

}
