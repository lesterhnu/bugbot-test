package goroutineleak

import (
	"fmt"
	"time"
)

// Run starts goroutines that can block forever, simulating a potential leak.
func Run() {
	ch := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("[goroutine_leak] worker %d waiting forever\n", id)
			<-ch // Bug pattern: no sender/close path, goroutine can never exit.
		}(i)
	}

	// Sleep briefly so leaked goroutines are visible in runtime behavior.
	time.Sleep(50 * time.Millisecond)
}
