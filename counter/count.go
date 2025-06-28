package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	display "github.com/AlexPassalis/command-line-applications-in-go/counter/display"
)

type Counts struct {
	bytes int
	words int
	lines int
}

func (c Counts) Add(other Counts) Counts {
	c.bytes += other.bytes
	c.words += other.words
	c.lines += other.lines

	return c
}

func (c Counts) Print(writer io.Writer, options display.Options, suffixes ...string) {
	stats := []string{}

	if options.ShouldShowHeader() {
		stats = append(stats, strconv.Itoa(c.lines))
		stats = append(stats, strconv.Itoa(c.words))
		stats = append(stats, strconv.Itoa(c.bytes))
	}

	if options.ShouldShowLines() {
		stats = append(stats, strconv.Itoa(c.lines))
	}

	if options.ShouldShowWords() {
		stats = append(stats, strconv.Itoa(c.words))
	}

	if options.ShouldShowBytes() {
		stats = append(stats, strconv.Itoa(c.bytes))
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

		res.bytes += size

		if rune == '\n' {
			res.lines++
		}

		isSpace := unicode.IsSpace(rune)
		if !isSpace && !isInsideWord {
			res.words++
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
