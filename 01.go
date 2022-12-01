package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AOC202201Helper(input string) ([]int, error) {
	var supplies []int

	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if line == "" {
			if sum != 0 {
				supplies = append(supplies, sum)
				sum = 0
			}
			continue
		}

		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, fmt.Errorf("can't parse \"%s\" as an integer: %v", line, err)
		}
		sum += lineInt
	}

	// Allow inputs without newlines at the end
	if sum != 0 {
		supplies = append(supplies, sum)
	}

	sort.Ints(supplies)

	return supplies, nil
}

func AOC202201SumMaxN(supplies []int, n int) (int, error) {
	if len(supplies) < n {
		return 0, fmt.Errorf("input %+v contains less than %d elements", supplies, n)
	}

	sum := 0
	for _, supply := range supplies[len(supplies)-n:] {
		sum += supply
	}

	return sum, nil
}

func AOC2022011(input string) (int, error) {
	supplies, err := AOC202201Helper(input)
	if err != nil {
		return 0, err
	}

	return AOC202201SumMaxN(supplies, 1)
}

func AOC2022012(input string) (int, error) {
	supplies, err := AOC202201Helper(input)
	if err != nil {
		return 0, err
	}

	return AOC202201SumMaxN(supplies, 3)
}
