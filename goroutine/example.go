package goroutine

import (
	"fmt"
	"time"
)

func Example() {
	go sayHello("Minh")

	sayHello("Thinh")
	time.Sleep(time.Second)
}

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello,%s\n", name)
	}
}
