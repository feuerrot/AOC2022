package main

import (
	"testing"
)

func TestAOC202208HiddenTrees(t *testing.T) {
	test := struct {
		input  string
		output int
	}{
		input: `30373
25512
65332
33549
35390`,
		output: 21,
	}

	forest, err := AOC202208ParseForest(test.input)
	if err != nil {
		t.Fatalf("AOC202208ParseForest() err: %v", err)
	}

	forest.HideTrees()
	got := forest.VisibleTrees()
	if got != test.output {
		t.Errorf("AOC202208Forest.VisibleTrees() missmatch:\nwant: %d\ngot:  %d", test.output, got)
	}
}
