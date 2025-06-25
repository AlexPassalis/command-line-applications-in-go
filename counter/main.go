package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0) // clears all log built-in prefixes.

	if len(os.Args) < 2 {
		log.Fatalln("error: no filename provided")
	}

	total := 0

	filenames := os.Args[1:]
	for _, filename := range filenames {
		wordCount := CountWordsInFile(filename)
		total += wordCount
		fmt.Println(wordCount, filename)
	}

	fmt.Println(total, "words")
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	return CountWords(file)
}

func CountWords(file io.Reader) int {
	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
