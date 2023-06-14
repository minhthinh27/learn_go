package defer_panic_recover_return

import "fmt"

func Run() {
	a := 1
	b := 2

	if (a + b) > 2 {
		panic("errors")
	}

	defer second()

	first()
}

func first() {
	fmt.Println("first")
}

func second() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("errors")

}
