package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var AOC202210CPUHalt = errors.New("HALT, STOPP")

type AOC202210CPU struct {
	Cycle int
	X     int
	Inst  []*AOC202210Instruction
}

func AOC202210NewCPU(instructions []*AOC202210Instruction) *AOC202210CPU {
	cpu := &AOC202210CPU{
		Cycle: 0,
		X:     1,
		Inst:  instructions,
	}

	return cpu
}

func (cpu *AOC202210CPU) Tick() (*int, error) {
	// cycle start
	cpu.Cycle += 1
	var signal *int

	if len(cpu.Inst) == 0 {
		return nil, AOC202210CPUHalt
	}
	inst := cpu.Inst[0]

	// cycle middle
	if (cpu.Cycle-20)%40 == 0 {
		sig := cpu.X * cpu.Cycle
		signal = &sig
	}

	if inst.Delay == 0 {
		cpu.X += inst.Value
		cpu.Inst = cpu.Inst[1:]
	} else {
		inst.Delay -= 1
	}

	return signal, nil
}

type AOC202210Instruction struct {
	Value int
	Delay int
}

func AOC202210ParseInstruction(input string) (*AOC202210Instruction, error) {
	data := strings.Split(input, " ")
	if len(data) > 2 {
		return nil, fmt.Errorf("can't parse %s as instruction", input)
	}

	if data[0] == "noop" {
		return &AOC202210Instruction{}, nil
	}

	value, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, fmt.Errorf("can't parse %s as integer: %v", data[1], err)
	}

	return &AOC202210Instruction{
		Value: value,
		Delay: 1,
	}, nil
}

func AOC202210ParseInstructions(input string) ([]*AOC202210Instruction, error) {
	lines := strings.Split(input, "\n")
	rtn := []*AOC202210Instruction{}
	for _, line := range lines {
		instruction, err := AOC202210ParseInstruction(line)
		if err != nil {
			return rtn, err
		}
		rtn = append(rtn, instruction)
	}

	return rtn, nil
}

func AOC2022101Helper(input string) (int, error) {
	instructions, err := AOC202210ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	cpu := AOC202210NewCPU(instructions)
	for {
		signal, err := cpu.Tick()
		if err == AOC202210CPUHalt {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("unexpected CPU error: %v", err)
		}
		if signal != nil {
			sum += *signal
		}
	}

	return sum, nil
}

func AOC2022101(input string) (string, error) {
	out, err := AOC2022101Helper(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", out), nil
}
