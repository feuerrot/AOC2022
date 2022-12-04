package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AOC202204Contains(line string) (bool, error) {
	parts := strings.Split(line, ",")
	lPart := strings.Split(parts[0], "-")
	rPart := strings.Split(parts[1], "-")

	left := []int{}
	right := []int{}

	for _, lValue := range lPart {
		v, err := strconv.Atoi(lValue)
		if err != nil {
			return false, fmt.Errorf("can't parse \"%s\" as integer: %v", lValue, err)
		}
		left = append(left, v)
	}
	for _, rValue := range rPart {
		v, err := strconv.Atoi(rValue)
		if err != nil {
			return false, fmt.Errorf("can't parse \"%s\" as integer: %v", rValue, err)
		}
		right = append(right, v)
	}

	if left[0] <= right[0] && left[1] >= right[1] {
		return true, nil
	}
	if right[0] <= left[0] && right[1] >= left[1] {
		return true, nil
	}
	return false, nil
}

func AOC2022041(input string) (int, error) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		contains, err := AOC202204Contains(line)
		if err != nil {
			return 0, err
		}
		if contains {
			sum += 1
		}
	}

	return sum, nil
}
