// Declaration: var m map[K]V
/* Properties
- Maps are reference types
*/
package main

import "fmt"

type User struct {
	Name string
}

func main() {
	var m map[string]int
	fmt.Println(m) // nil

	// Initialization
	// make function
	var make_map = make(map[string]int)

	fmt.Println(make_map) // map[]

	var literal_map = map[string]int{
		"a": 0,
		"b": 1,
	}

	fmt.Println(literal_map) // map[a:0 b:1]

	var user_map = map[string]User{
		"a": {"Peter"},
		"b": {"Seth"},
	}

	fmt.Println(user_map)

	// Adding values
	user_map["c"] = User{"Steve"}

	// Retrieving values
	fmt.Println("Key c:", user_map["c"])

	// If we use a key not present in the map, we get the zero value of the value type
	fmt.Println("Key d:", user_map["d"])

	// Retrieving a value also returns a boolean value that is true if the key exists and false otherwise
	d, ok := user_map["d"]
	fmt.Println("Key d:", d, ok)

	// Deleting a key
	delete(user_map, "c")
	fmt.Println(user_map)

	// Iteration
	for key, value := range user_map {
		fmt.Println("Key: %s, Value: %v", key, value)
	}
}
