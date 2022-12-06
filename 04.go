package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AOC202204FullOverlap(left, right []int) bool {
	// l: ..?456?..
	// r: ...456...
	if left[0] <= right[0] && left[1] >= right[1] {
		return true
	}
	// l: ...456...
	// r: ..?456?..
	if right[0] <= left[0] && right[1] >= left[1] {
		return true
	}

	return false
}

func AOC202204PartialOverlap(left, right []int) bool {
	if AOC202204FullOverlap(left, right) {
		return true
	}

	// l: ...456...
	// r: ...??6??.
	if left[1] >= right[0] && left[0] <= right[1] {
		return true
	}

	// l: ...??6??.
	// r: ...456...
	if right[1] >= left[0] && right[0] <= left[1] {
		return true
	}

	return false
}

func AOC202204Parse(line string) ([]int, []int, error) {
	parts := strings.Split(line, ",")
	lPart := strings.Split(parts[0], "-")
	rPart := strings.Split(parts[1], "-")

	left := []int{}
	right := []int{}

	for _, lValue := range lPart {
		v, err := strconv.Atoi(lValue)
		if err != nil {
			return nil, nil, fmt.Errorf("can't parse \"%s\" as integer: %v", lValue, err)
		}
		left = append(left, v)
	}
	for _, rValue := range rPart {
		v, err := strconv.Atoi(rValue)
		if err != nil {
			return nil, nil, fmt.Errorf("can't parse \"%s\" as integer: %v", rValue, err)
		}
		right = append(right, v)
	}

	return left, right, nil
}

func AOC2022041(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		left, right, err := AOC202204Parse(line)
		if err != nil {
			return "", err
		}

		contains := AOC202204FullOverlap(left, right)
		if contains {
			sum += 1
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func AOC2022042(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		left, right, err := AOC202204Parse(line)
		if err != nil {
			return "", err
		}

		contains := AOC202204PartialOverlap(left, right)
		if contains {
			sum += 1
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
