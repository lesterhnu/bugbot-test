package sliceoob

import "fmt"

// Run triggers an out-of-bounds access and recovers, so the app keeps running.
func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[slice_oob] recovered panic: %v\n", r)
		}
	}()

	arr := []int{10, 20, 30}
	idx := len(arr)

	// Bug pattern: arr[len(arr)] is always out of bounds.
	fmt.Printf("[slice_oob] value=%d\n", arr[idx])
}
