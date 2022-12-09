package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AOC202209Location struct {
	Head      []int
	Tail      []int
	Tailvisit map[string]interface{}
}

func absInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func AOC202209NewLocation() *AOC202209Location {
	return &AOC202209Location{
		Head:      []int{0, 0},
		Tail:      []int{0, 0},
		Tailvisit: make(map[string]interface{}),
	}
}

func (location AOC202209Location) Count() int {
	return len(location.Tailvisit)
}

func (location *AOC202209Location) Visit() {
	key := fmt.Sprintf("%d:%d", location.Tail[0], location.Tail[1])
	location.Tailvisit[key] = true
}

func (location *AOC202209Location) Move(movement AOC202209Instruction) {
	rdelta, udelta := movement.Dir()
	for i := 0; i < movement.Steps; i++ {
		headROld, headUOld := location.Head[0], location.Head[1]
		location.Head[0] += rdelta
		location.Head[1] += udelta

		rHeadTailDelta := location.Head[0] - location.Tail[0]
		uHeadTailDelta := location.Head[1] - location.Tail[1]
		if absInt(rHeadTailDelta) > 1 || absInt(uHeadTailDelta) > 1 {
			location.Tail = []int{headROld, headUOld}
		}
		location.Visit()
	}
}

type AOC202209Instruction struct {
	Direction string
	Steps     int
}

// Step calculates the R/U movement
func (instruction AOC202209Instruction) Dir() (int, int) {
	switch instruction.Direction {
	case "L":
		return -1, 0
	case "R":
		return 1, 0
	case "U":
		return 0, 1
	case "D":
		return 0, -1
	}

	return 0, 0
}

func AOC202209ParseInstuctions(input string) ([]AOC202209Instruction, error) {
	instructions := []AOC202209Instruction{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid command line: %s", line)
		}

		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("can't convert %s to int: %v", parts[1], err)
		}

		instructions = append(instructions, AOC202209Instruction{
			Direction: parts[0],
			Steps:     steps,
		})
	}

	return instructions, nil
}

func AOC2022091Helper(input string) (int, error) {
	instructions, err := AOC202209ParseInstuctions(input)
	if err != nil {
		return 0, fmt.Errorf("can't parse instructions: %v", err)
	}

	loc := AOC202209NewLocation()

	for _, instruction := range instructions {
		loc.Move(instruction)
	}

	return loc.Count(), nil
}

func AOC2022091(input string) (string, error) {
	data, err := AOC2022091Helper(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", data), nil
}
