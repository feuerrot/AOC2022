package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AOC202209Location struct {
	Node      [][]int
	Tailvisit map[string]interface{}
}

func absInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func sgnInt(in int) int {
	if in < 0 {
		return -1
	} else if in > 0 {
		return 1
	}
	return 0
}

func AOC202209NewLocation(length int) *AOC202209Location {
	rtn := &AOC202209Location{
		Node:      [][]int{},
		Tailvisit: make(map[string]interface{}),
	}
	for i := 0; i < length; i++ {
		rtn.Node = append(rtn.Node, []int{0, 0})
	}

	return rtn
}

func (location AOC202209Location) Count() int {
	return len(location.Tailvisit)
}

func (location *AOC202209Location) Visit() {
	key := fmt.Sprintf("%d:%d", location.Node[len(location.Node)-1][0], location.Node[len(location.Node)-1][1])
	location.Tailvisit[key] = true
}

func (location *AOC202209Location) Move(movement AOC202209Instruction) {
	for step := 0; step < movement.Steps; step++ {
		rdelta, udelta := movement.Dir()
		location.Node[0][0] += rdelta
		location.Node[0][1] += udelta
		for node := 1; node < len(location.Node); node++ {
			rHeadTailDelta := location.Node[node-1][0] - location.Node[node][0]
			uHeadTailDelta := location.Node[node-1][1] - location.Node[node][1]
			if absInt(rHeadTailDelta) > 1 || absInt(uHeadTailDelta) > 1 {
				rdelta = sgnInt(rHeadTailDelta)
				udelta = sgnInt(uHeadTailDelta)
			} else {
				rdelta = 0
				udelta = 0
			}

			location.Node[node][0] += rdelta
			location.Node[node][1] += udelta
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

	loc := AOC202209NewLocation(2)

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

func AOC2022092Helper(input string) (int, error) {
	instructions, err := AOC202209ParseInstuctions(input)
	if err != nil {
		return 0, fmt.Errorf("can't parse instructions: %v", err)
	}

	loc := AOC202209NewLocation(10)

	for _, instruction := range instructions {
		loc.Move(instruction)
	}

	return loc.Count(), nil
}

func AOC2022092(input string) (string, error) {
	data, err := AOC2022092Helper(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", data), nil
}
