package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type DisplayOptions struct {
	ShowBytes bool
	ShowWords bool
	ShowLines bool
}

func (d DisplayOptions) ShouldShowBytes() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowBytes
}

func (d DisplayOptions) ShouldShowWords() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowWords
}

func (d DisplayOptions) ShouldShowLines() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowLines
}

func main() {
	options := DisplayOptions{}

	flag.BoolVar(
		&options.ShowWords,
		"w",
		false,
		"Used to toggle whether or not to show the word count",
	)

	flag.BoolVar(
		&options.ShowBytes,
		"l",
		false,
		"Used to toggle whether or not to show the line count",
	)

	flag.BoolVar(
		&options.ShowLines,
		"c",
		false,
		"Used to toggle whether or not to show the byte count",
	)

	flag.Parse()

	log.SetFlags(0) // clears all log built-in prefixes.

	totals := Counts{}

	didError := false
	filenames := flag.Args()
	for _, filename := range filenames {
		counts, err := CountFile(filename)
		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		totals = totals.Add(counts)

		counts.Print(os.Stdout, options, filename)
	}

	if len(filenames) == 0 {
		GetCounts(os.Stdin).Print(os.Stdout, options)
	}

	if len(filenames) > 1 {
		totals.Print(os.Stdout, options, "total")
	}

	if didError {
		os.Exit(1)
	}
}
