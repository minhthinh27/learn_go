package main

import (
	"fmt"
	"sync"
	"time"
)

var delayForSwitchJob = time.Second * 3

func DoAJob(action string, duration time.Duration) {
	time.Sleep(delayForSwitchJob)
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), action)
	time.Sleep(duration)
}

var completeBoilingWater = false

func BoilingWaterJob(duration time.Duration) {
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Boil water")
	time.Sleep(duration)
	completeBoilingWater = true
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Ring! Ring! Ring! Completing boiling water ...")
}

func MakeASimpleCoffee() {
	// Async 3 job
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		BoilingWaterJob(time.Second * 60)
		//runtime.Gosched()
	}()

	go func() {
		defer wg.Done()
		DoAJob("Prepare coffee into the filter", time.Second*10)
		DoAJob("Press coffee", time.Second*10)
	}()
	wg.Wait()

	if completeBoilingWater {
		DoAJob("Put the boiled water into the filter", time.Second*5)
		DoAJob("Wait in some seconds ...", time.Second*60)
		fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Enjoy ...")
	}
}
