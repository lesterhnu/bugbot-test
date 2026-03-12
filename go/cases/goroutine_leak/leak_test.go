package goroutineleak

import (
	"runtime"
	"testing"
	"time"
)

func TestRunDoesNotLeakGoroutines(t *testing.T) {
	before := runtime.NumGoroutine()

	for i := 0; i < 20; i++ {
		Run()
	}

	// Give scheduler/GC a brief chance to finish cleanup.
	time.Sleep(100 * time.Millisecond)
	after := runtime.NumGoroutine()

	// Allow a small amount of runtime noise from the test harness.
	if after > before+2 {
		t.Fatalf("goroutines grew unexpectedly: before=%d after=%d", before, after)
	}
}
