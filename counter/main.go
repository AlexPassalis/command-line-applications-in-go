package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type DisplayOptions struct {
	ShowBytes  bool
	ShowWords  bool
	ShowLines  bool
	ShowHeader bool
}

func (d DisplayOptions) ShouldShowBytes() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowBytes
}

func (d DisplayOptions) ShouldShowWords() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowWords
}

func (d DisplayOptions) ShouldShowLines() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowLines
}

func (d DisplayOptions) ShouldShowHeader() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowHeader
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

	flag.BoolVar(
		&options.ShowHeader,
		"header",
		false,
		"Used to toggle whether or not to show the counts header",
	)

	flag.Parse()

	log.SetFlags(0) // clears all log built-in prefixes.

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := Counts{}

	if options.ShowHeader {
		fmt.Fprintln(os.Stdout, "lines\twords\tbytes")
	}

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
		GetCounts(os.Stdin).Print(writer, options)
	}

	if len(filenames) > 1 {
		totals.Print(os.Stdout, options, "total")
	}

	writer.Flush()

	if didError {
		os.Exit(1)
	}
}
