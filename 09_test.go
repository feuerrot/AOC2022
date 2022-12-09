package main

import "testing"

func TestAOC2022091(t *testing.T) {
	test := struct {
		input  string
		output int
	}{
		input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
		output: 13,
	}

	got, err := AOC2022091Helper(test.input)
	if err != nil {
		t.Fatalf("AOC2022091Helper() err: %v", err)
	}

	if got != test.output {
		t.Errorf("AOC2022091Helper() missmatch:\nwant: %d\ngot:  %d", test.output, got)
	}
}

func TestAOC202209LocationMove(t *testing.T) {
	test := struct {
		steps []AOC202209Instruction
		count int
	}{
		steps: []AOC202209Instruction{
			{
				Direction: "U",
				Steps:     2,
			},
			{
				Direction: "R",
				Steps:     2,
			},
		},
		count: 3,
	}

	loc := AOC202209NewLocation()
	for _, step := range test.steps {
		loc.Move(step)
	}

	if loc.Count() != test.count {
		t.Errorf("AOC202209Location.Count() missmatch after moves:\nwant: %d\ngot:  %d", test.count, loc.Count())
	}

}
