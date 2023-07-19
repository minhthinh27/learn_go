package excersie

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// the parent goroutine waits for the child
func WaitForResult() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("child : sent signal")
		ch <- 10
		//fmt.Println("child : sent signal")
	}()

	m := <-ch
	fmt.Println("parent : recv'd signal :", m)
	fmt.Println("ok")
	time.Sleep(time.Second)
}

// the parent goroutine creates 2000 child goroutines and waits for them to signal their results.
func FanOut() {
	children := 50
	ch := make(chan string, children)

	for i := 0; i < children; i++ {
		go func(c int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", c)
		}(i)
	}

	for children > 0 {
		m := <-ch
		children--
		fmt.Println(m)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
}

// In this pattern, the parent goroutine sends a signal to a child goroutine waiting to be told what to do.
func WaitForTask() {
	ch := make(chan string)

	go func() {
		m := <-ch
		fmt.Println("child : recv'd signal ", m)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	//fmt.Println("parent : sent signal")
	ch <- "data"
	fmt.Println("parent : sent signal")

	fmt.Println("ok")
	time.Sleep(time.Second)
}

//  the parent goroutine signals 100 pieces of work to a pool of child goroutines waiting for work to perform.
func Pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)

	for i := 0; i < g; i++ {
		go func(child int) {
			for work := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, work)
			}

			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(i)
	}

	const works = 100
	for w := 0; w < works; w++ {
		ch <- "work"
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
}

//  a semaphore is added to the fan out pattern to restrict the number of child goroutines that can be schedule to run.
func FanOutSem() {
	children := 50
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	fmt.Println(g)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
}

// crawlURL
const (
	numberOfURLs    = 10000
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	queue := make(chan int, 100)

	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Println("getURL")
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	return queue
}

func Start() {
	queue := startQueue()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Minute * 5)
}

// boundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signalling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.
func BoundedWorkPooling() {
	works := []string{"paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup

	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)

		}(i)
	}

	for _, work := range works {
		ch <- work
	}
	close(ch)

	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// drop: In this pattern, the parent goroutine signals 2000 pieces of work to
// a single child goroutine that can't handle all the work. If the parent
// performs a send and the child is not ready, that work is discarded and dropped.
func Drop() {
	work := 2000
	ch := make(chan string, 100)

	go func() {
		for m := range ch {
			fmt.Println("child : recv'd signal :", m)
		}
	}()

	for i := 0; i < work; i++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", i)
		default:
			fmt.Println("parent : dropped data :", i)

		}
	}
	close(ch)

	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// cancellation: In this pattern, the parent goroutine creates a child
// goroutine to perform some work. The parent goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the parent goroutine walks away.
func Cancellation() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case m := <-ch:
		fmt.Println("work complete", m)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
