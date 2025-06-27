package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	counter "github.com/AlexPassalis/command-line-applications-in-go"
	display "github.com/AlexPassalis/command-line-applications-in-go/display"
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

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := counter.Counts{}

	if arguments.ShowHeaders {
		fmt.Fprintln(os.Stdout, "lines\twords\tbytes")
	}

	didError := false
	filenames := flag.Args()
	for _, filename := range filenames {
		counts, err := counter.CountFile(filename)
		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		totals = totals.Add(counts)

		counts.Print(os.Stdout, options, filename)
	}

	if len(filenames) == 0 {
		counter.GetCounts(os.Stdin).Print(writer, options)
	}

	if len(filenames) > 1 {
		totals.Print(os.Stdout, options, "total")
	}

	writer.Flush()

	if didError {
		os.Exit(1)
	}
}
