package goroutine

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 100
}

func deposit(val int, wg *sync.WaitGroup) {
	mutex.Lock()
	balance += val
	mutex.Unlock()
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mutex.Lock()
	balance -= val
	mutex.Unlock()
	wg.Done()
}

func Example5() {
	var wg sync.WaitGroup
	wg.Add(3)
	go deposit(20, &wg)
	go withdraw(80, &wg)
	go deposit(40, &wg)
	wg.Wait()
	fmt.Printf("balance is: %d\n", balance)
}
