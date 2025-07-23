package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for index := range 10 {
			channel <- index
 		}
	}()

	for value := range channel {
		fmt.Println(value)
	}
}
