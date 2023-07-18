package channels

import "fmt"

func Close() {
	ch := make(chan int, 10)
	var a string

	go func() {
		a = "ok"
		close(ch)
	}()

	<-ch
	fmt.Println(a)
}

func Close1() {
	ch := make(chan int, 10)
	var a string

	go func() {
		a = "ok"
		close(ch)
	}()

	ch <- 10
	fmt.Println(a)
}

func Close2() {
	ch := make(chan int)
	var a string

	go func() {
		a = "ok"
		close(ch)
	}()

	ch <- 10
	fmt.Println(a)
}
