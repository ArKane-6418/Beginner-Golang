// Generics means parameterized types and allow programmers to write code where the type can be specified later since it's not immediately relevant
/*
func fnName[T constraint]() {
    ...
}
T is the type parameter and constraint will be the interface that allows any type implementing the interface
*/

/* When to use generics
- Functions that operate on arrays, slices, maps, and channels
- General purpose data structures like stack or linked list
- To reduce code duplication
*/

package main

// Instead of defining our own constraint, we can use the constraints package
import (
	"fmt"
	"golang.org/x/exp/constraints
)

// Can define our own custom constraint to allow for + operations

// type SumConstraint interface {
// 	int | float64 | string
// }

// In Go 1.18, we can replace interface{} with any

func sum[T constraints.Ordered](a, b T) T {
	return a + b
}

/*
func sumInt(a, b int) int {
    return a + b
}

func sumFloat(a, b float64) float64 {
    return a + b
}

func sumString(a, b string) string {
    return a + b
}
*/

func main() {
	// Note how aside from the types, the sum functions are identical
	fmt.Println(sum[int](1, 2))
	fmt.Println(sum[float64](4.0, 2.0))
	fmt.Println(sum[string]("a", "b"))

	// Go 1.18 also comes with type inference so we don't need to pass type arguments
	fmt.Println(sum(1, 2))
	fmt.Println(sum(4.0, 2.0))
	fmt.Println(sum("a", "b"))
}
