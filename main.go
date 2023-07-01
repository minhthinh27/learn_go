package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chlBoyA := make(chan string, 1)
	chlBoyB := make(chan string, 1)
	ctx := context.Background()

	go play(chlBoyA, chlBoyB, ctx)

	chlBoyB <- "pong"

	time.Sleep(time.Millisecond)
}

func play(chlA, chlB chan string, ctx context.Context) {
	for {
		select {
		case <-chlA:
			fmt.Println("ping")
			chlB <- "pong"
		case <-chlB:
			fmt.Println("pong")
			chlA <- "ping"
		case <-ctx.Done():
			return
		}
	}
}
