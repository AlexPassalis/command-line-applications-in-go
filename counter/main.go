package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0) // clears all log built-in prefixes.

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

	if len(filenames) == 0 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if didError {
		os.Exit(1)
	}
}
