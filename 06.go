package main

import "fmt"

func AOC202206FindPreamble(input string) int {
	for i := 0; i < len(input)-4; i++ {
		if input[i] == input[i+1] ||
			input[i] == input[i+2] ||
			input[i] == input[i+3] ||
			input[i+1] == input[i+2] ||
			input[i+1] == input[i+3] ||
			input[i+2] == input[i+3] {
			continue
		}
		return i + 4
	}
	return 0
}

func AOC2022061(input string) (string, error) {
	index := AOC202206FindPreamble(input)
	return fmt.Sprintf("%d", index), nil
}
