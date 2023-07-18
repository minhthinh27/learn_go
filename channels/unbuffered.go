package channels

import (
	"fmt"
	"sync"
	"time"
)

func Unbuffered() {
	ch := make(chan int)
	var a string

	go func() {
		fmt.Println("channel reciver data")
		<-ch
		a = "ok"
	}()

	fmt.Println("channel send data")
	ch <- 10

	fmt.Println("continued processing")
	fmt.Println(a)
}

func Unbuffered1() {
	ch := make(chan int)
	var a string

	go func() {
		fmt.Println("channel send data")
		ch <- 10
		a = "ok"
	}()

	fmt.Println("channel reciver data")
	time.Sleep(time.Second * 3)
	<-ch

	fmt.Println("continued processing")
	fmt.Println(a)
}

func Unbuffered2() {
	ch1 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch1 <- "Hello TGA"
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("this is message :%s \n", <-ch1)
	}()

	wg.Wait()
}

func Unbuffered3() {
	ch := make(chan string)
	go func() {
		fmt.Println("send reciver")
		ch <- "Golang"
		fmt.Println("sended reciver")
	}()

	fmt.Println("wait reciver")
	fmt.Printf("%s\n", <-ch)
	fmt.Println("recivered")

	time.Sleep(time.Second * 3)
}

func Unbuffered4() {
	ch := make(chan string)

	select {
	case ch <- "hello":
		fmt.Println("sent message")
	default:
		fmt.Println("no message sent")
	}

}
