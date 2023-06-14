package closure

import "fmt"

// https://www.calhoun.io/what-is-a-closure/#closuresprovidedataisolation
func Run() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

func newCounter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
