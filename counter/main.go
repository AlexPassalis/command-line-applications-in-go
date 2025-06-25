package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "./words.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
