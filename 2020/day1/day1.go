package day1

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

const SUM = 2020

func Solve(data io.Reader) (int, error) {
	expenses, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	return solve(expenses, SUM), nil
}

func solve(expenses []int, max int) int {
	set := make(map[int]bool)
	for _, expense := range expenses {
		goal := max - expense
		if _, ok := set[goal]; ok {
			return goal * expense
		} else {
			set[expense] = true
		}
	}

	return 0
}

func SolveTriple(data io.Reader) (int, error) {
	expenses, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	sort.IntSlice(expenses).Sort()

	for _, expense := range expenses {
		goal := SUM - expense
		i := sort.SearchInts(expenses, goal)
		res := solve(expenses[:i], goal)

		if res != 0 {
			return expense * res, nil
		}
	}

	return 0, nil
}

func parseInput(data io.Reader) ([]int, error) {
	res := []int{}

	s := bufio.NewScanner(data)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, nil
}
