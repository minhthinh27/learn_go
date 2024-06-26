package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

func Example1() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addOne(&wg)
	}
	wg.Wait()
	fmt.Println("Giá trị của x là: ", x)
}

func addOne(wg *sync.WaitGroup) {
	atomic.AddInt64(&x, 1)
	wg.Done()
}
