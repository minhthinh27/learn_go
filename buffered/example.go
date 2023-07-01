package buffered

import (
	"fmt"
	"time"
)

func Run() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	go worker("thinh", jobs, results)
	go worker("thinh1", jobs, results)
	go worker("thinh2", jobs, results)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func example1() {
	bufferChanel := make(chan int, 10)

	bufferChanel <- 1
	bufferChanel <- 2
	bufferChanel <- 3

	close(bufferChanel)

	// /* Receive with for-range */
	// for value := range bufferChanel {
	// 	fmt.Printf("value in channel: %v \n", value)
	// }

	for true {
		value, found := <-bufferChanel
		if found {
			fmt.Printf("value in channel: %v \n", value)
		} else {
			break
		}
	}
}

func worker(id string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished  job", j)

		results <- j * 2

	}
}
