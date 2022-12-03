package main

import (
	"strings"
	"testing"
)

func TestAOC202202(t *testing.T) {
	var tests = []struct {
		input  string
		points int
		score  int
		target string
	}{
		{
			input:  "A Y",
			points: 2,
			score:  6,
			target: "X",
		},
		{
			input:  "B X",
			points: 1,
			score:  0,
			target: "X",
		},
		{
			input:  "C Z",
			points: 3,
			score:  3,
			target: "X",
		},
	}

	for _, test := range tests {
		inputParts := strings.Split(test.input, " ")

		got := AOC202202Win(inputParts[0], inputParts[1])
		if test.score != got {
			t.Errorf("AOC202202Win() missmatch:\nwant: %d\ngot:  %d", test.score, got)
		}

		got = AOC202202Points(inputParts[1])
		if test.points != got {
			t.Errorf("AOC202202Points() missmatch:\nwant: %s->%d\ngot:  %s->%d", inputParts[0], test.points, inputParts[1], got)
		}

		target := AOC202202Lookup(inputParts[0], inputParts[1])
		if test.target != target {
			t.Errorf("AOC202202Points() missmatch:\nwant: %s->%s\ngot:  %s->%s", inputParts[0], test.target, inputParts[1], target)
		}
	}
}
