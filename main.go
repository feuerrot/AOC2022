package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Task func(string) (string, error)

func main() {
	tasks := map[int][]Task{
		1:  {AOC2022011, AOC2022012},
		2:  {AOC2022021, AOC2022022},
		3:  {AOC2022031, AOC2022032},
		4:  {AOC2022041, AOC2022042},
		5:  {AOC2022051, AOC2022052},
		6:  {AOC2022061, AOC2022062},
		7:  {AOC2022071, AOC2022072},
		8:  {AOC2022081, AOC2022082},
		9:  {AOC2022091, AOC2022092},
		10: {AOC2022101},
	}
	days := []int{}
	for day := range tasks {
		days = append(days, day)
	}
	sort.Ints(days)

	for _, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input/2022%02d", day))
		if err != nil {
			log.Fatalf("can't open input: %v", err)
		}

		for i, task := range tasks[day] {
			res, err := task(string(input))
			if err != nil {
				log.Fatalf("error in Task %d/%d: %v", day, i+1, err)
			}
			log.Printf("AOC2022 %02d/%d: %s", day, i+1, res)
		}
	}
}
