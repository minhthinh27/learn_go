package excersie

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// the parent goroutine waits for the child
func WaitForResult() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- 10
		fmt.Println("child : sent signal")
	}()

	m := <-ch
	fmt.Println("parent : recv'd signal :", m)

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
	fmt.Println("parent : sent signal")
	ch <- "data"

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
