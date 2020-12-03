package tobTraj_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/hennersz/advent-of-code/2020/tobTraj"
)

func Test_solve(t *testing.T) {
	in, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatalf("Error occurred opening test file: %v", err)
	}
	inBuffer, err := ioutil.ReadAll(in)
	if err != nil {
		t.Fatalf("Error occurred opening test file: %v", err)
	}

	cases := []struct {
		Description string
		Input       io.Reader
		Slopes      io.Reader
		Want        int
	}{
		{
			"case1",
			bytes.NewReader(inBuffer),
			bytes.NewReader([]byte(strings.Join([]string{"3 1"}, "\n"))),
			7,
		},
		{
			"case2",
			bytes.NewReader(inBuffer),
			bytes.NewReader([]byte(strings.Join([]string{"1 1", "3 1", "5 1", "7 1", "1 2"}, "\n"))),
			336,
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := tobTraj.Solve(test.Input, test.Slopes)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
