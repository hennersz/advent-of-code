package reportRepair_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/hennersz/advent-of-code/2020/reportRepair"
)

func Test_solve(t *testing.T) {
	cases := []struct {
		Description string
		Input       io.Reader
		Want        int
	}{
		{
			"case1",
			bytes.NewReader([]byte(strings.Join([]string{"1721", "979", "366", "299", "675", "1456"}, "\n"))),
			514579,
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := reportRepair.Solve(test.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func Test_solveN(t *testing.T) {
	cases := []struct {
		Description string
		Input       io.Reader
		N           int
		Want        int
	}{
		{
			"case1",
			bytes.NewReader([]byte(strings.Join([]string{"1721", "979", "366", "299", "675", "1456"}, "\n"))),
			3,
			241861950,
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := reportRepair.SolveTriple(test.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
