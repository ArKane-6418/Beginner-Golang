// A struct is a user-defined type that contains a collection of named fields. Basically, It is used to group related data together to form a single unit.
// Defined like so: type Person struct {}

/* Properties
- Structs are value types
- EMpty structs occupy zero bytes of storage
*/
package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

// Embedding (not preferred)
// type Superhero struct {
// 	Person
// 	Alias string
// }

// Composition (preferred)
type Superhero struct {
	Person Person
	Alias  string
}

func main() {

	// Can use the new keyword
	p := new(Person)
	p.FirstName = "Tatsuya"
	p.LastName = "Shiba"
	p.Age = 18
	fmt.Println("Person", p)
	// Declaration
	var p1 Person

	fmt.Println("Person 1:", p1) // Zero value: Person 1: {  0}
	fmt.Printf("Person 1 FirstName: %s, LastName: %s, Age: %d\n", p1.FirstName, p1.LastName, p1.Age)

	var p2 = Person{FirstName: "Joshua", LastName: "Ong", Age: 21}

	fmt.Println("Person 2:", p2)

	// Can also do pointers to structs, and Go is smart enough to auto-dereference the pointer
	ptr := &p2

	fmt.Println((*ptr).FirstName)
	fmt.Println(ptr.FirstName)
	// We can choose to initialize a subset of fields where the uninitialized fields take their type's zero value
	var p3 = Person{
		FirstName: "Tony",
		LastName:  "Stark",
	}

	fmt.Println("Person 3:", p3)

	// We can also omit field name but we need to provide all values otherwise it will fail
	var p4 = Person{"Yuichi", "Katagiri", 17}
	fmt.Println("Person 4:", p4)

	// Can also declare an anonymous struct
	var a = struct {
		Name string
	}{"Golang"}

	fmt.Println("Anonymous:", a)

	superperson := Person{"Clark", "Kent", 40}
	s := Superhero{superperson, "Superman"}
	// Embedding approach
	// s := Superhero{}

	// s.FirstName = "Clark"
	// s.LastName = "Kent"
	// s.Age = 40
	// s.Alias = "Superman"

	fmt.Println(s)
}
