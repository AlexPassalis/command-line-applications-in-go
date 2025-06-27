package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Counts struct {
	Bytes int
	Words int
	Lines int
}

func (c Counts) Add(other Counts) Counts {
	c.Bytes += other.Bytes
	c.Words += other.Words
	c.Lines += other.Lines

	return c
}

func (c Counts) Print(writer io.Writer, options DisplayOptions, suffixes ...string) {
	stats := []string{}

	if options.ShouldShowHeader() {
		stats = append(stats, strconv.Itoa(c.Lines))
		stats = append(stats, strconv.Itoa(c.Words))
		stats = append(stats, strconv.Itoa(c.Bytes))
	}

	if options.ShouldShowLines() {
		stats = append(stats, strconv.Itoa(c.Lines))
	}

	if options.ShouldShowWords() {
		stats = append(stats, strconv.Itoa(c.Words))
	}

	if options.ShouldShowBytes() {
		stats = append(stats, strconv.Itoa(c.Bytes))
	}

	line := strings.Join(stats, "\t") + "\t"
	fmt.Fprint(writer, line)
	suffixStr := strings.Join(suffixes, " ")
	if suffixStr != "" {
		fmt.Fprintf(writer, " %s", suffixStr)
	}

	fmt.Fprint(writer, "\n")
}

func GetCounts(file io.Reader) Counts {
	res := Counts{}

	isInsideWord := false
	reader := bufio.NewReader(file)

	for {
		rune, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		res.Bytes += size

		if rune == '\n' {
			res.Lines++
		}

		isSpace := unicode.IsSpace(rune)
		if !isSpace && !isInsideWord {
			res.Words++
		}

		isInsideWord = !isSpace
	}

	return res
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	counts := GetCounts(file)

	return counts, nil
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
		rune, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if rune == '\n' {
			linesCount++
		}

	}

	return linesCount
}

func CountBytes(reader io.Reader) int {
	byteCount, _ := io.Copy(io.Discard, reader)
	return int(byteCount)
}
