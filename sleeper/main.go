package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		DoSomethingSlowly("a")
	}()

	go func() {
		defer waitGroup.Done()
		DoSomethingSlowly("b")
	}()

	waitGroup.Wait() // wait here for both goroutines
}

func DoSomethingSlowly(name string) {
	fmt.Println("Doing something slowly:", name)
	time.Sleep(time.Second * 5)
	fmt.Println("Finished:", name)
}
