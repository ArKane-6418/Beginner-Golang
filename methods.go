// Go is not an OOP language
// A method is nothing but a function with a special receiver argument. Let's see how we can declare methods
// func (variable T) Name(params) (returnTypes) {}
// The receiver argument has a name and a type. It appears between the func keyword and the method name

/* Properties
- Go is smart enough to interpret our function call correctly and hence, pointer receiver method calls are just syntactic sugar provided by Go for convenience. (&c).UpdateName(...)
- We can omit the variable part of the receiver as well if we're not using it: func (Car) UpdateName(...) {}
- Methods are not limited to structs but can also be used with non-struct types as well
*/

package main

import "fmt"

type Car struct {
	Name string
	Year int
}

// Accessing the instance of Car using the receiver variable c is like the 'this' keyword in OOP
func (c Car) IsLatest() bool {
	return c.Year >= 2017
}

// Methods with pointer receivers

func (c *Car) UpdateName(name string) {
	c.Name = name
}

func main() {
	c := Car{"Tesla", 2021}
	c.UpdateName("Toyota")
	fmt.Println("IsLatest", c.IsLatest())
	fmt.Println("Car:", c)
}
