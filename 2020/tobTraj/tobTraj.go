package tobTraj

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

const TREE = '#'

type slope struct {
	across, down int
}

func Solve(input io.Reader, slopesInput io.Reader) (int, error) {
	slopesScanner := bufio.NewScanner(slopesInput)
	var slopes []*slope

	for slopesScanner.Scan() {
		parts := strings.Split(slopesScanner.Text(), " ")

		across, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}

		down, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		slopes = append(slopes, &slope{across, down})
	}

	inputBuffer, err := ioutil.ReadAll(input)
	if err != nil {
		return 0, err
	}

	total := 1
	for _, s := range slopes {
		res, err := solve(bytes.NewReader(inputBuffer), s.across, s.down)
		if err != nil {
			return 0, err
		}
		total *= res
	}

	return total, nil
}

func solve(input io.Reader, across, down int) (int, error) {
	r := bufio.NewScanner(input)
	count := 0
	lineIndex := 0
	LINE_LENGTH := 0

	for r.Scan() {
		if LINE_LENGTH == 0 {
			LINE_LENGTH = len(r.Text())
		}

		pos := r.Text()[lineIndex]
		if pos == TREE {
			count++
		}

		lineIndex += across
		if lineIndex >= LINE_LENGTH {
			lineIndex = lineIndex % LINE_LENGTH
		}

		for i := down - 1; i > 0; i-- {
			r.Scan()
		}
	}

	return count, nil
}
