package goroutine

import "fmt"

func Example3() {
	c := make(chan string)
	go func() {
		c <- "tga"
	}()
	get := <-c
	fmt.Println(get)
}

func Example3V2() {
	c := make(chan string)
	go func() {
		get := <-c
		fmt.Println(get)
	}()

	c <- "tga_v2"
}
