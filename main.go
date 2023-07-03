package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	eventChannel := make(chan int)
	wg := new(sync.WaitGroup)

	go func() {
		for i := 0; i <= 100; i++ {

			wg.Add(1)
			go action(wg, i, eventChannel)
			wg.Wait()

		}

		close(eventChannel)
	}()

	for v := range eventChannel {
		fmt.Println(v)
	}

	time.Sleep(time.Second * 3)
}

func action(wg *sync.WaitGroup, i int, chn chan<- int) {
	defer wg.Done()
	chn <- i
}
