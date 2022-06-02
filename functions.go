package main

import "fmt"

func main() {
	myFunction("Hello")
	fmt.Println(myNextFunction("This is p1", "This is p2"))
	add := myClosureFunction()
	add(5)
	fmt.Println(add(10))
	sum := addMultiple(1, 2, 3, 5)
	fmt.Println(sum)

	// defer (lets us postpone the execution of a function until the surrounding function returns)
	// defer executes in a stack, so the bottom println gets executed before the one above it
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")
	fmt.Println("Doing some work...")
}

func myFunction(p1 string) {
	fmt.Println(p1)
}

// Shorthand if consecutive parameters have the same type
func myNextFunction(p1, p2 string) string {
	msg := fmt.Sprintf("p1: %s, p2: %s", p1, p2)
	return msg
}

// Multiple return values
func myMultiValueFunction(p1 string) (string, int) {
	msg := fmt.Sprintf("%s function", p1)
	return msg, 10
}

// Named returns
func myNamedFunction(p1 string) (s string, i int) {
	s = fmt.Sprintf("%s function", p1)
	i = 10

	return
}

// Functions as values
func myValueFunction() {
	fn := func() {
		fmt.Println("inside fn")
	}

	fn()
}

// Closures are lexically scoped, which means functions can access the values in scope when defining the function.
// This is a function that returns a function that takes an int and returns an int
func myClosureFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}

// Variadic Functions (take zero or multiple arguments using the ... ellipses operator)
func addMultiple(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}
