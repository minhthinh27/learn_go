package main

import (
	"learn_go/closure"
	"learn_go/difference"
	"learn_go/goroutine"
	"time"
)

func main() {
	closure.Run()
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
