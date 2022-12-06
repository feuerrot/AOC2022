package main

import (
	"fmt"
	"strings"
)

func AOC202203Priority(input rune) int {
	// a->z =>  1->26
	// A->Z => 27->52
	if 'a' <= input && input <= 'z' {
		return int(input - 'a' + 1)
	}

	return int(input - 'A' + 26 + 1)
}

func AOC202203Item(line string) rune {
	partlen := len(line) / 2

	// O(n^2), yeah!
	for _, lRune := range line[:partlen] {
		for _, rRune := range line[partlen:] {
			if lRune == rRune {
				return lRune
			}
		}
	}

	return ' '
}

func AOC202203Multiitem(first, second, third string) rune {
	// O(n^3), yeah!
	for _, lRune := range first {
		for _, mRune := range second {
			for _, rRune := range third {
				if lRune == mRune && mRune == rRune {
					return lRune
				}
			}
		}
	}

	return ' '
}

func AOC202203Round(line string) int {
	item := AOC202203Item(line)
	prio := AOC202203Priority(item)
	return prio
}

func AOC2022031(input string) (string, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		prio := AOC202203Round(line)
		sum += prio
	}

	return fmt.Sprintf("%d", sum), nil
}

func AOC2022032(input string) (string, error) {
	sum := 0
	lines := strings.Split(input, "\n")
	rounds := len(lines) / 3
	for round := 0; round < rounds; round++ {
		item := AOC202203Multiitem(lines[round*3], lines[round*3+1], lines[round*3+2])
		prio := AOC202203Priority(item)
		sum += prio
	}

	return fmt.Sprintf("%d", sum), nil
}
