// The fork-join model essentially implies that a child process splits from its parent process to run concurrently with the parent process. After completing its execution, the child process merges back into the parent process. The point where it joins back is called the join point.
// Trickiest part about using goroutines is when they will stop
// Important to know that goroutines run in the same address space, so access to shared memory must be synchronized
package main

import "fmt"

func speak(arg string) {
	fmt.Println(arg)
}

func main() {
	go speak("Hello world")
}
