package main

import (
	"learn_go/difference"
	"learn_go/goroutine"
	"learn_go/useful_closure"
	"time"
)

func main() {
	useful_closure.Run()
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
