package excersie

import (
	"fmt"
	"time"
)

func GamePingPong() {
	boyA := make(chan string, 1)
	boyB := make(chan string, 1)
	quit := make(chan string, 1)

	go func(a, b chan string) {
		for {
			select {
			case <-a:
				fmt.Println("ping")
				time.Sleep(time.Millisecond)
				b <- "pong"
			case <-b:
				fmt.Println("pong")
				time.Sleep(time.Millisecond)
				a <- "ping"
			case <-quit:
				fmt.Println("end game")
				return
			default:
				fmt.Println("preparing play game")
			}
		}
	}(boyA, boyB)

	boyB <- "pong"

	// for no limit
	//select {}

	time.Sleep(time.Second * 2)
}
