package reportRepair

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

	set := make(map[int]bool)
	for _, expense := range expenses {
		goal := SUM - expense
		if _, ok := set[goal]; ok {
			return goal * expense, nil
		} else {
			set[expense] = true
		}
	}

	return 0, nil
}

func solve(expenses []int, max int) (int, error) {
	set := make(map[int]bool)
	for _, expense := range expenses {
		goal := max - expense
		if _, ok := set[goal]; ok {
			return goal * expense, nil
		} else {
			set[expense] = true
		}
	}

	return 0, nil
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
		res, err := solve(expenses[:i], goal)
		if err != nil {
			return 0, err
		}

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
