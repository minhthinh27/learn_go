package unbuffered

import (
	"fmt"
	"time"
)

// channel unbuffered send and receive value

func Example2() {
	myChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("myChan send", i)
			myChan <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-myChan)
		fmt.Println("myChan received", i)
	}
}
