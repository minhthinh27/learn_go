package unbuffered

import "fmt"

func Run() {
	chanelAtMainRoutine := make(chan int)
	go func() {
		fmt.Println("Start getting value from channel ...")
		<-chanelAtMainRoutine
		fmt.Println("Complete getting value from channel ...")

	}()
	chanelAtMainRoutine <- 5
	fmt.Println("Completing ...")
}
