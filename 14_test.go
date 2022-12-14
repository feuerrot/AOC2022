package main

import "testing"

func TestAOC202214Part1(t *testing.T) {
	test := struct {
		input  string
		output int
	}{
		input: `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
		output: 24,
	}

	got, err := AOC2022141Helper(test.input)
	if err != nil {
		t.Fatalf("AOC202214Helper() err: %v", err)
	}

	if got != test.output {
		t.Errorf("AOC202214Helper() missmatch:\nwant: %d\ngot:  %d", test.output, got)
	}
}
