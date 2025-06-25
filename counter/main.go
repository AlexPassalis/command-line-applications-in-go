package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "./words.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	PrintFileContents(file)

	/*
		wordCount := CountWords(file)
		fmt.Println(wordCount)
	*/
}

func PrintFileContents(file *os.File) {
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)

	totalSize := 0
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		totalSize += size

		fmt.Print(string(buffer[:size]))
	}

	fmt.Println("total bytes read:", totalSize)
}

func CountWords(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
