package excersie

import (
	"fmt"
	"time"
)

func GamePingPong() {
	boyA := make(chan string, 1)
	boyB := make(chan string, 1)

	go func(a, b chan string) {
		for {
			select {
			case m := <-a:
				fmt.Println(m)
				time.Sleep(time.Millisecond)
				b <- "pong"
			case m := <-b:
				fmt.Println(m)
				time.Sleep(time.Millisecond)
				a <- "ping"
			default:
				fmt.Println("fuck")
			}
		}
	}(boyA, boyB)

	boyB <- "pong"

	// for no limit
	select {}
}
