package anonymous_functions

import "fmt"

var DoSuff func() = func() {}

/*
The big difference is that we could assign a new function to the DoStuff variable at runtime,
allowing us to dynamically change what DoStuff() does.

The only real difference between a regular function and an anonymous one is that anonymous functions aren’t declared at a package level.
*/
func Run() {
	DoSuff()

	// action
	DoSuff = func() { fmt.Println("action") }
	DoSuff()

	// another action
	DoSuff = func() { fmt.Println("another action") }
	DoSuff()
}
