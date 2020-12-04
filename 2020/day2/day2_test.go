package day2_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/hennersz/advent-of-code/2020/day2"
)

func Test_SolveRange(t *testing.T) {
	cases := []struct {
		Description string
		Input       io.Reader
		Want        int
	}{
		{
			"case1",
			bytes.NewReader([]byte(strings.Join([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, "\n"))),
			2,
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := day2.SolveRange(test.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func Test_SolvePosition(t *testing.T) {
	cases := []struct {
		Description string
		Input       io.Reader
		Want        int
	}{
		{
			"case1",
			bytes.NewReader([]byte(strings.Join([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, "\n"))),
			1,
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := day2.SolvePosition(test.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
