package main

import (
	"fmt"
	"log"
	"os"
)

type Task func(string) (int, error)

func main() {
	tasks := map[int][]Task{
		1: {AOC2022011, AOC2022012},
		2: {AOC2022021, AOC2022022},
		3: {AOC2022031, AOC2022032},
		4: {AOC2022041, AOC2022042},
		5: {AOC2022051, AOC2022052},
	}

	for day := range tasks {
		input, err := os.ReadFile(fmt.Sprintf("input/2022%02d", day))
		if err != nil {
			log.Fatalf("can't open input: %v", err)
		}

		for i, task := range tasks[day] {
			res, err := task(string(input))
			if err != nil {
				log.Fatalf("error in Task %d/%d: %v", day, i+1, err)
			}
			log.Printf("AOC2022 %02d/%d: %d", day, i+1, res)
		}
	}
}
