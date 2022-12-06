package main

import (
	"fmt"
)

func AOC202206FindUnique(input string, length int) int {
OUTER:
	for i := 0; i < len(input)-length; i++ {
		for first := i; first < i+length-1; first++ {
			for second := first + 1; second < i+length; second++ {
				if input[first] == input[second] {
					continue OUTER
				}
			}
		}
		return i + length
	}

	return 0
}

func AOC202206FindUniqueJump(input string, length int) int {
OUTER:
	for i := 0; i < len(input)-length; i++ {
		for first := i; first < i+length-1; first++ {
			for second := first + 1; second < i+length; second++ {
				if input[first] == input[second] {
					i = first
					continue OUTER
				}
			}
		}
		return i + length
	}

	return 0
}

func AOC202206FindUniqueReverseJump(input string, length int) int {
OUTER:
	for i := length - 1; i < len(input)-1; i++ {
		for first := i; first > i-length+1; first-- {
			for second := first - 1; second > i-length; second-- {
				if input[first] == input[second] {
					i = second + length - 1
					continue OUTER
				}
			}
		}
		return i + 1
	}

	return 0
}

func AOC2022061(input string) (string, error) {
	index := AOC202206FindUniqueReverseJump(input, 4)
	return fmt.Sprintf("%d", index), nil
}

func AOC2022062(input string) (string, error) {
	index := AOC202206FindUniqueReverseJump(input, 14)
	return fmt.Sprintf("%d", index), nil
}
