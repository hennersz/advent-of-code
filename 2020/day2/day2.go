package day2

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type entry struct {
	min, max int
	required rune
	password string
}

func SolveRange(data io.Reader) (int, error) {
	entries, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, e := range entries {
		if isValidRange(e) {
			count++
		}
	}
	return count, nil
}

func SolvePosition(data io.Reader) (int, error) {
	entries, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, e := range entries {
		if isValidPosition(e) {
			count++
		}
	}
	return count, nil
}

func parseInput(data io.Reader) ([]*entry, error) {
	res := []*entry{}

	s := bufio.NewScanner(data)
	for s.Scan() {
		re := regexp.MustCompile(`(\d+)-(\d+)\s+([a-z]+):\s+([a-z]+)`)
		parts := re.FindStringSubmatch(s.Text())

		min, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		max, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		e := &entry{
			min,
			max,
			[]rune(parts[3])[0],
			parts[4],
		}
		res = append(res, e)
	}

	return res, nil
}

func isValidRange(e *entry) bool {
	count := 0

	for _, char := range e.password {
		if char == e.required {
			count++
		}
	}

	return count >= e.min && count <= e.max
}

func isValidPosition(e *entry) bool {
	firstPosition := e.password[e.min-1] == byte(e.required)
	secondPosition := e.password[e.max-1] == byte(e.required)

	return firstPosition != secondPosition
}
