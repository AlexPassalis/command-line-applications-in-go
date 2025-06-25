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

	didError := false
	filenames := os.Args[1:]
	for _, filename := range filenames {
		wordCount, err := CountWordsInFile(filename)
		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}
		total += wordCount
		fmt.Println(wordCount, filename)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if didError {
		os.Exit(1)
	}
}

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	return CountWords(file), nil
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
