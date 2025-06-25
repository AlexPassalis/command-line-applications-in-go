package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("Please enter your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatalln("failed to scan name", err)
	}

	var powerlevel uint
	fmt.Printf("Hello %s. What is your powerlevel: ", name)
	_, err = fmt.Scanln(&powerlevel)
	if err != nil {
		log.Fatalln("failed to scan powerlevel", err)
	}

	fmt.Printf("%s's powerlevel: %d\n", name, powerlevel)

	if powerlevel > 9000 {
		fmt.Println("It's over 9000 !!!")
	}
}
