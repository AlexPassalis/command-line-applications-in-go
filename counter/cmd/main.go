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

	channel := make(chan FileCountsResult)

	didError := false
	for _, filename := range filenames {
		go func() {
			defer waitGroup.Done()

			counts, err := counter.CountFile(filename)
			if err != nil {
				didError = true
				fmt.Fprintln(os.Stderr, "counter:", err)
				return
			}

			channel <- FileCountsResult{
				counts:   counts,
				filename: filename,
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for result := range channel {
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
