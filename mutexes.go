// Mutex prevents other processes from entering a critical section of data while a process occupies it to prevent race conditions from happening.
// A critical section can be a piece of code that must not be run by multiple threads at once because the code contains shared resources
/* We can use Mutex using the following functions:
- Lock() acquires or holds the lock
- Unlock() releases the lock
- TryLock() tries to lock and reports whether it succeeded
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()
}
