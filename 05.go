package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type AOC202205Crate []rune
type AOC202205Crates []AOC202205Crate

func AOC202205ParseCrates(input string) (AOC202205Crates, error) {
	rtn := AOC202205Crates{}

	lines := strings.Split(input, "\n")
	maxStack := len(lines)
	boxLine := lines[maxStack-1]
	for i := 0; i < len(boxLine); i++ {
		if boxLine[i] == ' ' {
			continue
		}
		rtn = append(rtn, AOC202205Crate{})
	}

	for row := maxStack - 2; row >= 0; row-- {
		for column, char := range lines[row] {
			if 'A' <= char && char <= 'Z' {
				col := (column - 1) / 4
				rtn[col] = append(rtn[col], char)
			}
		}
	}

	return rtn, nil
}

func (crates AOC202205Crates) AOC202205ApplyCommands(commands []AOC202205Command) (AOC202205Crates, error) {
	for _, cmd := range commands {
		if len(crates[cmd.src-1]) < cmd.count {
			return crates, fmt.Errorf("can't move %d crates from %d to %d: %d has only %d elements", cmd.count, cmd.src, cmd.dst, cmd.src, len(crates[cmd.src-1]))
		}
		for i := 0; i < cmd.count; i++ {
			sLen := len(crates[cmd.src-1])
			sData := crates[cmd.src-1][sLen-1]
			crates[cmd.src-1] = crates[cmd.src-1][:sLen-1]
			crates[cmd.dst-1] = append(crates[cmd.dst-1], sData)
		}
	}

	return crates, nil
}

func (crates AOC202205Crates) AOC202205ApplyCommands2(commands []AOC202205Command) (AOC202205Crates, error) {
	for _, cmd := range commands {
		if len(crates[cmd.src-1]) < cmd.count {
			return crates, fmt.Errorf("can't move %d crates from %d to %d: %d has only %d elements", cmd.count, cmd.src, cmd.dst, cmd.src, len(crates[cmd.src-1]))
		}
		sLen := len(crates[cmd.src-1])
		sData := crates[cmd.src-1][sLen-cmd.count:]
		crates[cmd.src-1] = crates[cmd.src-1][:sLen-cmd.count]
		for _, char := range sData {
			crates[cmd.dst-1] = append(crates[cmd.dst-1], char)
		}
	}

	return crates, nil
}

func (crates AOC202205Crates) AOC202205GetTop() string {
	var out string
	for _, crate := range crates {
		out += string(crate[len(crate)-1])
	}

	return out
}

type AOC202205Command struct {
	src   int
	dst   int
	count int
}

func AOC202205ParseCommands(input string) ([]AOC202205Command, error) {
	rtn := []AOC202205Command{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		src, err := strconv.Atoi(parts[3])
		if err != nil {
			return rtn, fmt.Errorf("can't convert %s to int: %v", parts[3], err)
		}
		dst, err := strconv.Atoi(parts[5])
		if err != nil {
			return rtn, fmt.Errorf("can't convert %s to int: %v", parts[5], err)
		}
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return rtn, fmt.Errorf("can't convert %s to int: %v", parts[1], err)
		}

		rtn = append(rtn, AOC202205Command{
			src:   src,
			dst:   dst,
			count: count,
		})
	}

	return rtn, nil
}

func AOC2022051Wrapper(input string) (string, error) {
	stateCommand := strings.Split(input, "\n\n")
	crates, err := AOC202205ParseCrates(stateCommand[0])
	if err != nil {
		return "", fmt.Errorf("can't parse crates: %v", err)
	}

	commands, err := AOC202205ParseCommands(stateCommand[1])
	if err != nil {
		return "", fmt.Errorf("can't parse commands: %v", err)
	}

	crates, err = crates.AOC202205ApplyCommands(commands)
	if err != nil {
		return "", fmt.Errorf("can't apply commands: %v", err)
	}

	return crates.AOC202205GetTop(), nil
}

func AOC2022051(input string) (int, error) {
	out, err := AOC2022051Wrapper(input)
	if err != nil {
		return 0, err
	}
	log.Printf("AOC2022 05/1: %s", out)

	return 0, nil
}

func AOC2022052Wrapper(input string) (string, error) {
	stateCommand := strings.Split(input, "\n\n")
	crates, err := AOC202205ParseCrates(stateCommand[0])
	if err != nil {
		return "", fmt.Errorf("can't parse crates: %v", err)
	}

	commands, err := AOC202205ParseCommands(stateCommand[1])
	if err != nil {
		return "", fmt.Errorf("can't parse commands: %v", err)
	}

	crates, err = crates.AOC202205ApplyCommands2(commands)
	if err != nil {
		return "", fmt.Errorf("can't apply commands: %v", err)
	}

	return crates.AOC202205GetTop(), nil
}

func AOC2022052(input string) (int, error) {
	out, err := AOC2022052Wrapper(input)
	if err != nil {
		return 0, err
	}
	log.Printf("AOC2022 05/2: %s", out)

	return 0, nil
}
