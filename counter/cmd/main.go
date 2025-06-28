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
		fmt.Fprintln(os.Stdout, "lines\twords\tbytes")
	}

	filenames := flag.Args()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(filenames))

	lock := sync.Mutex{}

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

			lock.Lock()
			defer lock.Unlock()
			// State mutated below
			totals = totals.Add(counts)
			counts.Print(tabWriter, options, filename)

		}()
	}

	waitGroup.Wait()

	if len(filenames) > 1 {
		totals.Print(os.Stdout, options, "total")
	}

	if len(filenames) == 0 {
		counter.GetCounts(os.Stdin).Print(tabWriter, options)
	}

	tabWriter.Flush()

	if didError {
		os.Exit(1)
	}
}
