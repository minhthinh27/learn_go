package goroutine

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func Example2() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func(str string) {
		defer wg.Done()
		for i := 0; i < 50000; i++ {
			num := strconv.Itoa(i)
			sum := sha1.Sum([]byte(num))

			fmt.Printf("%s: %05d: %x\n", str, i, sum)
		}
	}("A")

	go func(str string) {
		defer wg.Done()
		for i := 0; i < 50000; i++ {
			num := strconv.Itoa(i)
			sum := sha1.Sum([]byte(num))

			fmt.Printf("%s: %05d: %x\n", str, i, sum)
		}
	}("B")

	fmt.Println("Wait to finish")
	wg.Wait()

}
