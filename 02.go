package main

import (
	"fmt"
	"strings"
)

func AOC202202Win(opponent, self string) int {
	var points = map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	return points[opponent][self]
}

func AOC202202Points(self string) int {
	var points = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	return points[self]
}

func AOC202202Round(opponent, self string) (int, error) {
	return AOC202202Win(opponent, self) + AOC202202Points(self), nil
}

func AOC2022021(input string) (int, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		points, err := AOC202202Round(parts[0], parts[1])
		if err != nil {
			return 0, fmt.Errorf("Line: %s: %v", line, err)
		}

		sum += points
	}

	return sum, nil
}

func AOC202202Lookup(opponent, target string) string {
	var lookup = map[string]map[string]string{
		"A": {
			"X": "Z",
			"Y": "X",
			"Z": "Y",
		},
		"B": {
			"X": "X",
			"Y": "Y",
			"Z": "Z",
		},
		"C": {
			"X": "Y",
			"Y": "Z",
			"Z": "X",
		},
	}

	return lookup[opponent][target]
}

func AOC2022022(input string) (int, error) {

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		points, err := AOC202202Round(parts[0], AOC202202Lookup(parts[0], parts[1]))
		if err != nil {
			return 0, fmt.Errorf("Line: %s: %v", line, err)
		}

		sum += points
	}

	return sum, nil
}
