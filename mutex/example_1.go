package mutex

import (
	"fmt"
	"sync"
)

var x int64 = 0
var mutex = &sync.Mutex{}

func Example1() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOne(&wg, mutex)
	}
	wg.Wait()
	fmt.Println("Giá trị của x là: ", x)
}

func addOne(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	defer mutex.Unlock()
	wg.Done()
}
