package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	if len(data) == 0 { // guard clause to prevent the edge case of an empty document
		return 0
	}

	wordDetected := false
	wordCount := 0

	const spaceCharRune = ' ' // or const spaceCharASCII = 32 where 32 is the ASCII value of the Space character

	for _, x := range data {
		if x == spaceCharRune {
			wordCount++ // increase the value of wordCount by 1, every time x is a space character (incorrect algorithm for many edge cases)
		} else {
			wordDetected = true
		}
	}

	if !wordDetected { // if only spaces (0 words detected), return 0
		return 0
	}

	wordCount++

	return wordCount
}
