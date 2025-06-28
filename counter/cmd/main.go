package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"

	counter "github.com/AlexPassalis/command-line-applications-in-go/counter"
	display "github.com/AlexPassalis/command-line-applications-in-go/counter/display"
)

type FileCountsResult struct {
	counts   counter.Counts
	filename string
	err      error
}

func main() {
	arguments := display.NewOptionsArguments{}

	flag.BoolVar(
		&arguments.ShowWords,
		"w",
		false,
		"Used to toggle whether or not to show the word count",
	)

	flag.BoolVar(
		&arguments.ShowLines,
		"l",
		false,
		"Used to toggle whether or not to show the line count",
	)

	flag.BoolVar(
		&arguments.ShowBytes,
		"c",
		false,
		"Used to toggle whether or not to show the byte count",
	)

	flag.BoolVar(
		&arguments.ShowHeaders,
		"header",
		false,
		"Used to toggle whether or not to show the counts header",
	)

	flag.Parse()

	options := display.NewOptions(arguments)

	log.SetFlags(0) // clears all log built-in prefixes.

	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := counter.Counts{}

	if arguments.ShowHeaders {
		fmt.Fprintln(tabWriter, "lines\twords\tbytes")
	}

	filenames := flag.Args()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(filenames))

	didError := false

	channel := CountFiles(filenames)

	for result := range channel {
		if result.err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", result.err)
			continue
		}

		totals = totals.Add(result.counts)
		result.counts.Print(tabWriter, options, result.filename)
	}

	if len(filenames) > 1 {
		totals.Print(tabWriter, options, "total")
	}

	if len(filenames) == 0 {
		counter.GetCounts(os.Stdin).Print(tabWriter, options)
	}

	tabWriter.Flush()

	if didError {
		os.Exit(1)
	}
}

func CountFiles(filenames []string) <-chan FileCountsResult { // typeof receive only channel
	channel := make(chan FileCountsResult)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(filenames))

	for _, filename := range filenames {
		go func() {
			defer waitGroup.Done()

			result, err := CountsForFile(filename)
			channel <- FileCountsResult{
				counts:   result,
				filename: filename,
				err:      err,
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	return channel
}

func CountsForFile(filename string) (counter.Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return counter.Counts{}, err
	}
	defer file.Close()

	return counter.GetCounts(file), nil
}
