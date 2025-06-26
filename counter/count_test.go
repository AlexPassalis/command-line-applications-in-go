package main_test

import (
	"strings"
	"testing"

	counter "github.com/AlexPassalis/command-line-applications-in-go"
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
			result := counter.CountWords(r)
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
			result := counter.CountLines(r)
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
			result := counter.CountBytes(r)
			if result != tc.wants {
				t.Logf("expected: %d got %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}
