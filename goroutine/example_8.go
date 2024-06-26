package goroutine

import "fmt"

func Example8() {
	myChan := make(chan int)

	go receiveAndSend(myChan)
	myChan <- 1

	fmt.Printf("value received: %d", <-myChan)
}
func receiveAndSend(c chan int) {
	fmt.Printf("Receive: %d\n", <-c)
	fmt.Printf("Sending 2....")
	c <- 2
}
