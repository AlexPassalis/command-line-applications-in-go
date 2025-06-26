package main

import (
	"bufio"
	"io"
	"os"
)

type counts struct {
	bytes int
	words int
	lines int
}

func CountFile(filename string) (counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return counts{}, err
	}
	defer file.Close()

	lineCount := CountLines(file)
	byteCount := CountBytes(file)
	wordCount := CountWords(file)

	return counts{
		bytes: byteCount,
		words: wordCount,
		lines: lineCount,
	}, nil
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

func CountLines(r io.Reader) int {
	linesCount := 0

	reader := bufio.NewReader(r)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			linesCount++
		}

	}

	return linesCount
}

func CountBytes(r io.Reader) int {
	byteCount, _ := io.Copy(io.Discard, r)
	return int(byteCount)
}
