package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AOC2022011(input string) (int, error) {
	var supplies []int

	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if line == "" {
			supplies = append(supplies, sum)
			sum = 0
			continue
		}

		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("can't parse \"%s\" as an integer: %v", line, err)
		}
		sum += lineInt
	}

	max := 0
	for _, value := range supplies {
		if value > max {
			max = value
		}
	}

	return max, nil
}
