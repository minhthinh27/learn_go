package closure

import "fmt"

func Run() {
	example3()
}

type add func(a int, b int) int

func example() {
	var a add = func(a, b int) int {
		return a + b
	}

	s := a(7, 7)
	fmt.Println(s)
}

func add2(a func(a, b int) int) {
	fmt.Println(a(27, 27))
}

func example2() {
	f := func(a, b int) int {
		return a + b
	}

	add2(f)
}

func add3() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}

	return f
}

func example3() {
	s := add3()
	fmt.Println(s(27, 27))
}
