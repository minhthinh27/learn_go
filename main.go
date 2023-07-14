package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// var key = "the quick brown fox jumps over the lazy dog"
	// var message = "vkbs bs t suepuv"
	// var nums = []int{0, 1, 4, 6, 7, 10}
	// var diff = 3
	//var rings = "B0B6G0R6R0R6G9"
	//var words = []string{"cd", "ac", "dc", "ca", "zz"}
	//fmt.Println(countPoints(rings))
	//waitForResult()
	drop()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		fmt.Println("child send data")
		time.Sleep(time.Second * 2)
		ch <- "data"
	}()

	d := <-ch

	fmt.Printf("parent reciver data: %s", d)

	time.Sleep(time.Second * 3)
}

func fanOut() {
	children := 2000

	ch := make(chan string, children)

	for i := 0; i <= children; i++ {
		go func(child int) {
			time.Sleep(time.Millisecond * 2)
			ch <- fmt.Sprintf("children: %d send data", child)
		}(i)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Printf("reciver data: %s", d)
	}

	time.Sleep(time.Second * 5)
}

func waitForTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("children reciver: ", d)
	}()

	time.Sleep(time.Second * 2)

	ch <- "data"
	fmt.Println("parent send data")

	time.Sleep(time.Second * 3)
}

func pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Println("children reciver word", child, d)
			}

			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	work := 20
	for i := 0; i <= work; i++ {
		ch <- "data parents"
		fmt.Println("parent send word", i)
	}

	close(ch)

	time.Sleep(time.Second * 3)
}

func fanOutSem() {
	children := 20
	ch := make(chan string, children)
	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				//fmt.Println("child : sent signal :", child)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("parent revicer singal of children: ", children)
		children--
	}

	time.Sleep(time.Second)
}

func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 20: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}

			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
}

func drop() {
	const cap = 100

	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	works := 1000
	for i := 0; i < works; i++ {
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
}
