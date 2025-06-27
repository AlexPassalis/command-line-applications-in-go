package count

import (
	"bytes"
	"strings"
	"testing"

	display "github.com/AlexPassalis/command-line-applications-in-go/display"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentense.  This is another.",
			wants: 7,
		},
		{
			name:  "prefixed multiple spaces",
			input: "  Hello",
			wants: 1,
		},
		{
			name:  "suffixed multiple spaces",
			input: "Hello  ",
			wants: 1,
		},
		{
			name:  "tab character in code",
			input: "Hello\tWord\n",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := GetCounts(r).words
			if result != tc.wants {
				t.Logf("expected: %d got %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words, 1 new line",
			input: "one two three four five \n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new line at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi newline string",
			input: "\n\n\n\n\n",
			wants: 5,
		},
		{
			name:  "multi word line string",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := GetCounts(r).lines
			if result != tc.wants {
				t.Logf("expected: %d got %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "     ",
			wants: 5,
		},
		{
			name:  "newlines and words",
			input: "one\ntwo\nthree\nfour\t\n",
			wants: 20,
		},
		{
			name:  "unicode characters",
			input: "Ѷѯ",
			wants: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := GetCounts(r).bytes
			if result != tc.wants {
				t.Logf("expected: %d got %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestGetCounts(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants Counts
	}{
		{
			name:  "simple five words",
			input: "one two three four five\n",
			wants: Counts{
				lines: 1,
				words: 5,
				bytes: 24,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := GetCounts(r)
			if result != tc.wants {
				t.Logf("expected: %d got %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestPrintCounts(t *testing.T) {
	type inputs struct {
		counts   Counts
		options  display.NewOptionsArguments
		filename []string
	}

	testCases := []struct {
		name  string
		input inputs
		wants string
	}{
		{
			name: "simple five words.txt",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
				options: display.NewOptionsArguments{
					ShowLines: true,
					ShowWords: true,
					ShowBytes: true,
				},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "simple five words.txt show lines",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
				options: display.NewOptionsArguments{
					ShowLines: true,
					ShowWords: false,
					ShowBytes: false,
				},
			},
			wants: "1\t words.txt\n",
		},
		{
			name: "simple five words.txt no options",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
		{
			name: "simple five words.txt show words",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
				options: display.NewOptionsArguments{
					ShowLines: false,
					ShowWords: true,
					ShowBytes: false,
				},
			},
			wants: "5\t words.txt\n",
		},
		{
			name: "simple five words.txt show bytes and lines",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
				options: display.NewOptionsArguments{
					ShowLines: true,
					ShowWords: false,
					ShowBytes: true,
				},
			},
			wants: "1\t24\t words.txt\n",
		},
		{
			name: "empty filename",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				filename: []string{"words.txt"},
				options: display.NewOptionsArguments{
					ShowLines: true,
					ShowWords: true,
					ShowBytes: true,
				},
			},
			wants: "1\t5\t24\t words.txt\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			tc.input.counts.Print(buffer, display.NewOptions(tc.input.options), tc.input.filename...)
			if buffer.String() != tc.wants {
				t.Logf("expected: %s got %s", tc.wants, buffer.String())
				t.Fail()
			}
		})
	}
}

func TestAddCounts(t *testing.T) {
	type inputs struct {
		counts Counts
		other  Counts
	}

	testCases := []struct {
		name  string
		input inputs
		wants Counts
	}{
		{
			name: "simple add by one",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 24,
				},
				other: Counts{
					lines: 1,
					words: 1,
					bytes: 1,
				},
			},
			wants: Counts{
				lines: 2,
				words: 6,
				bytes: 25,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			totals := tc.input.counts
			res := totals.Add(tc.input.other)
			if res != tc.wants {
				t.Logf("expected: %v got %v", tc.wants, totals)
				t.Fail()
			}
		})
	}
}
