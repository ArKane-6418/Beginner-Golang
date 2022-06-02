// A channel is a communications pipe between goroutines
// Declaring a channel: var ch chan T
/* Properties
- A send to a nil channel blocks forever
	var c chan string
	c <- "Hello, World!" // Panic: all goroutines are asleep - deadlock!

- A receive from a nil channel blocks forever
	var c chan string
	fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!

- A send to a closed channel panics
	var c = make(chan string, 1)
	c <- "Hello, World!"
	close(c)
	c <- "Hello, Panic!" // Panic: send on closed channel

- A receive from a closed channel returns the zero value immediately
	var c = make(chan int, 2)
	c <- 5
	c <- 4
	close(c)
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", <-c) // Output: 5 4 0 0
	}
*/

package main

import "fmt"

// Sending and receiving data

func speak(arg string, ch chan string) {
	// channel <- data
	ch <- arg // Send
}

func main() {
	var c chan string

	fmt.Println(c) // <nil>

	// Can initialize using make
	ch := make(chan string)
	fmt.Println(ch)

	go speak("Hello World", ch)

	// Receivers can test whether a channel has been closed by assigning a second param to receive expression
	data, ok := <-ch // Receive data
	fmt.Println(data, ok)

	// Buffered channels accept a limited number of values without a corresponding receiver for them
	// Buffer length (capacity) can be specified using the second argument of make

	buffer_ch := make(chan string, 2)
	go speak("Hello World", buffer_ch)
	go speak("Hi again", buffer_ch)

	data1 := <-buffer_ch
	fmt.Println(data1)

	data2 := <-buffer_ch
	fmt.Println(data2)

	// By default a channel is unbuffered and has a capacity of 0, hence we omit the second argument to the make function.

	// Directional channels
	// We can specify if a channel is meant to only send or receive values
	// chan<- for send only (send to channel)
	// <-chan for receive only (receive from channel)

	// Need to close channel
	close(ch)
	close(buffer_ch)

	// We can iterate over values received from a channel

	new_ch := make(chan string, 2)
	go speak("Bye World", buffer_ch)
	go speak("Bye again", buffer_ch)

	close(new_ch)
	for data := range new_ch {
		fmt.Println(data)
	}
}
