package channels

import (
	"fmt"
	"sync"
)

func Buffered() {
	ss := make(chan string, 2)

	ss <- "Golang"
	ss <- "Nodejs"

	fmt.Println(<-ss)
	fmt.Println(<-ss)

}

func Buffered1() {
	ch1 := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		ch1 <- "Hi,TGA"
		ch1 <- "Nice to meet you"
	}()

	go func() {
		defer wg.Done()

		fmt.Printf("this is message 1 %s\n", <-ch1)
		fmt.Printf("this is message 2 %s\n", <-ch1)
	}()

	wg.Wait()
}
