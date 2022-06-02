// An interface in Go is an abstract type that is defined using a set of method signatures
// The interface defines the behavior for similar types of objects.

/* Properties
- Zero value is nil
- Can embed interfaces like structs
	type interface1 interface {
    	Method1()
	}
	type interface2 interface {
		Method2()
	}
	type interface3 interface {
		interface1
		interface2
	}

- Interface values are comparable
- An interface value can be thought of as a tuple consisting of a value and concrete type
package main

import "fmt"

// Empty interface allows us to handle values of unknown types
var x interface{}

type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}

type socket struct{}

// We have defined a Plug method that accepts a mobile device
// We don't want to define this same method for the other 4 devices
// Therefore, we define a PowerDrawer interface which expects the Draw method

type PowerDrawer interface {
	Draw(power int)
}

func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)

	// Typer assertion provides access to an interface value's underlying concrete value
	var i interface{} = "hello"

	// Typer assertion returns the underlying value and a boolean that reports whether the assertion succeeded
	i2, ok := i.(string)
	fmt.Println(i2, ok)

	// Type Switch is used to determine the type of a variable of type empty interface{}
	var p interface{}
	p = "hello"

	switch p := p.(type) {
	case string:
		fmt.Printf("string: %s\n", p)
	case bool:
		fmt.Printf("boolean: %v\n", p)
	case int:
		fmt.Printf("integer: %d\n", p)
	default:
		fmt.Printf("unexpected: %T\n", p)
	}
}
