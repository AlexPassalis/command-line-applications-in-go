package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	filename := "./words.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWordsInFile(file)
	fmt.Println(wordCount)

}

func CountWordsInFile(file *os.File) int {
	wordCount := 0
	isInsideWord := false

	const bufferSize = 8192
	buffer := make([]byte, bufferSize)

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		isInsideWord = !unicode.IsSpace(rune(buffer[size-1])) && isInsideWord

		bufferCount := CountWords(buffer[:size])
		if isInsideWord {
			bufferCount -= 1
		}

		wordCount += bufferCount

		isInsideWord = !unicode.IsSpace(rune(buffer[size-1]))
	}

	return wordCount

}

func CountWords(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
