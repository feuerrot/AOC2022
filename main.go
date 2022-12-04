package main

import (
	"log"
	"os"
)

type Task func(string) (int, error)

func main() {
	days := []struct {
		day   int
		input string
		tasks []Task
	}{
		{
			day:   1,
			input: "input/202201",
			tasks: []Task{
				AOC2022011,
				AOC2022012,
			},
		}, {
			day:   2,
			input: "input/202202",
			tasks: []Task{
				AOC2022021,
				AOC2022022,
			},
		}, {
			day:   3,
			input: "input/202203",
			tasks: []Task{
				AOC2022031,
				AOC2022032,
			},
		}, {
			day:   4,
			input: "input/202204",
			tasks: []Task{
				AOC2022041,
				AOC2022042,
			},
		},
	}

	for _, day := range days {
		input, err := os.ReadFile(day.input)
		if err != nil {
			log.Fatalf("can't open input: %v", err)
		}

		for i, task := range day.tasks {
			res, err := task(string(input))
			if err != nil {
				log.Fatalf("error in Task %d/%d: %v", day.day, i+1, err)
			}
			log.Printf("AOC2022 %02d/%d: %d", day.day, i+1, res)
		}
	}
}
