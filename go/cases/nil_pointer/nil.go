package nilpointer

import "fmt"

type user struct {
	Name string
}

// Run triggers a nil pointer dereference and recovers for demo continuity.
func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[nil_pointer] recovered panic: %v\n", r)
		}
	}()

	var u *user

	// Bug pattern: dereferencing nil pointer.
	fmt.Printf("[nil_pointer] user=%s\n", u.Name)
}
