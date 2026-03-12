package goroutineleak

import (
	"fmt"
	"sync"
)

// Run starts workers and guarantees they can all exit.
func Run() {
	ch := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("[goroutine_leak] worker %d waiting for shutdown signal\n", id)
			<-ch
		}(i)
	}

	close(ch)
	wg.Wait()
}
