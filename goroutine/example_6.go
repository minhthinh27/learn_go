package goroutine

import (
	"fmt"
	"sync"
)

func init() {
	balance = 100
}

func depositV2(val int, wg *sync.WaitGroup, ch chan bool, time string) {
	ch <- true
	balance += val
	test := <-ch
	fmt.Println("depositV2", time)
	fmt.Println("depositV2", test)
	wg.Done()
}

func withdrawV2(val int, wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	balance -= val
	test := <-ch
	fmt.Println("withdrawV2", test)
	wg.Done()
}

func Example6() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	wg.Add(3)
	go depositV2(10, &wg, ch, "1")
	go withdrawV2(80, &wg, ch)
	go depositV2(40, &wg, ch, "2")
	wg.Wait()
	fmt.Printf("balance is: %d\n", balance)
}
