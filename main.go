package main

import (
	"learn_go/defer_panic_recover_return"
	"learn_go/difference"
	"learn_go/goroutine"
	"time"
)

func main() {
	defer_panic_recover_return.Run()
}

func actionGoroutine() {
	// ex
	go goroutine.MakeASimpleCoffee()

	/* Wait a cup of coffee is made*/
	time.Sleep(time.Second * 150)
}

func actionDifference1() {
	difference.Difference1()
}
