// WaitGroup helps us wait for multiple goroutines to finish
/*
- Add(int) function takes in an integer value which is essentially the number of goroutines that the waitgroup has to wait for. This function must be called before we execute a goroutine.
- Done() function is called within the goroutine to signal that the goroutine has successfully executed.
- Wait() function blocks the program until all the goroutines specified by Add() have invoked Done() from within.
*/

package main

import (
	"fmt"
	"sync"
)

// func work() {
//     fmt.Println("working...")
// }

// Can also pass waitgroup directly
func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup

	// go func() {
	//     defer wg.Done()
	//     work()
	// }()

	wg.Add(4)

	go work(&wg)
	go work(&wg)
	go work(&wg)
	go work(&wg)

	wg.Wait()

	// WaitGroups must not be copied, if it's explicitly passed into functions, should be done via pointer
}
