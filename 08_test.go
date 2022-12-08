package main

import (
	"testing"
)

func TestAOC202208(t *testing.T) {
	test := struct {
		input  string
		hidden int
		score  int
	}{
		input: `30373
25512
65332
33549
35390`,
		hidden: 21,
		score:  8,
	}

	forest, err := AOC202208ParseForest(test.input)
	if err != nil {
		t.Fatalf("AOC202208ParseForest() err: %v", err)
	}

	forest.HideTrees()
	got := forest.VisibleTrees()
	if got != test.hidden {
		t.Errorf("AOC202208Forest.VisibleTrees() missmatch:\nwant: %d\ngot:  %d", test.hidden, got)
	}

	forest.CalculateScore()
	score := forest.Highscore()
	if score != test.score {
		t.Errorf("AOC202208Forest.Highscore() missmatch:\nwant: %d\ngot:  %d", test.score, score)
	}
}
