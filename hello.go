package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// Initializing variables
	var foo string = "Go is awesome"

	fmt.Println(foo)

	// Multi-initialization
	var foo2, bar string = "Hello", "World"
	fmt.Println(foo2, bar)

	// Alternatively
	var (
		a string = "Hello"
		b string = "World"
	)
	fmt.Println(a, b)

	// Type can be inferred
	foo3 := "Shorthand"
	fmt.Println(foo3)

	// Constants
	const constant = "This is a constant"
	fmt.Println(constant)

	// Bool
	var value bool = false
	var isItTrue bool = true
	fmt.Println(value, isItTrue)

	// Numeric types
	var i int = 404                     // Platform dependent
	var i8 int8 = 127                   // Signed 8-bit integer (-128 to 127)
	var i16 int16 = 32767               // Signed 16-bit integer (-2^15 to 2^15 - 1)
	var i32 int32 = -2147483647         // Signed 32-bit integer (-2^31 to 2^31 - 1)
	var i64 int64 = 9223372036854775807 // Signed 64-bit integer (-2^63 to 2^63 - 1)

	fmt.Println(i, i8, i16, i32, i64)

	var ui uint = 404                     // Platform dependent
	var ui8 uint8 = 255                   // Unsigned 8-bit integer (0 to 255)
	var ui16 uint16 = 65535               // Unsigned 16-bit integer (0 to 2^16 - 1)
	var ui32 uint32 = 2147483647          // Unsigned 32-bit integer (0 to 2^32 - 1)
	var ui64 uint64 = 9223372036854775807 // Unsigned 64-bit integer (0 to 2^64 - 1)

	fmt.Println(ui, ui8, ui16, ui32, ui64)

	// Integer alias types
	type byte = uint8
	type rune = int32
	var by byte = 'a'
	var ru rune = '🍕'

	fmt.Println(by, ru)

	// FLoating point (default is float64)
	var f32 float32 = 1.7812 // IEEE-754 32-bit
	var f64 float64 = 3.1415 // IEEE-754 64-bit

	fmt.Println(f32, f64)
	// Complex
	var c1 complex128 = complex(10, 1) // Both real and imaginary parts are 64-bit
	var c2 complex64 = 12 + 4i         // Both real and imaginary parts are 32-bit

	fmt.Println(c1, c2)

	// Zero values
	var in int
	var fl float64
	var bo bool
	var st string
	fmt.Printf("%v %v %v %q\n", in, fl, bo, st) // 0 0 false ""

	// Type conversion
	i_42 := 42
	f_42 := float64(i_42)
	u_42 := uint(f_42)
	fmt.Printf("%T %T %T\n", i_42, f_42, u_42)

	// Alias types (allow developers to provide an alternate name for an existing type and use it interchangeably with the underlying type)
	type MyAlias = string
	var str MyAlias = "I am an alias"
	fmt.Printf("%T - %s\n", str, str)

	// Defined types (the defined type is different from its underlying type so it cannot be used interchangeably)
	type MyDefined string
	var str2 MyDefined = "I am defined"
	fmt.Printf("%T - %s\n", str2, str2) // Output: main.MyDefined

	// String formatting
	// Sprintf, Sprintf, Sprintln return the string instead of printing it
	s := fmt.Sprintf("hex:%x bin:%b", 10, 10)
	fmt.Println(s)

	// Multiline
	msg := `
	Hello from
	multiline
	`

	fmt.Println(msg)

	// if/else
	x := 10

	if x > 5 {
		fmt.Println("x is gt 5")
	} else if x > 10 {
		fmt.Println("x is gt 10")
	} else {
		fmt.Println("else case")
	}

	// compact if
	if y := 2; y < 3 {
		fmt.Println("y is less than 3")
	}

	// switch case only runs the first case whose value is equal to the condition expression (break is automatically added)
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("time to work!")
	case "friday":
		fmt.Println("let's party")
	default:
		fmt.Println("browse memes")
	}

	// Can also do shorthand declaration

	switch month := "august"; month {
	case "january":
		fmt.Println("Too cold")
	case "april":
		fmt.Println("Getting warmer")
	case "august":
		fmt.Println("Summer!")
		fallthrough
	// Use fallthrough keyword to transfer control to the next case
	default:
		fmt.Println("A month of the year")
	}

	// For loop
	for j := 0; j < 10; j++ {
		if j < 2 {
			continue
		}

		fmt.Println(j)

		if j > 5 {
			break
		}
	}

	fmt.Println("We broke out!")

	// Init and post statements are optional so we can write a for loop that behaves like a while loop
	k := 0

	for k < 10 {
		k += 1
	}

	// No condition means infinite loop

}
