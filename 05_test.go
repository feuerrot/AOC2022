package main

import "testing"

func TestAOC2022051(t *testing.T) {
	test := struct {
		input  string
		output string
	}{
		input: `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
		output: "CMZ",
	}

	got, err := AOC2022051Wrapper(test.input)
	if err != nil {
		t.Fatalf("AOC2022051() err: %v", err)
	}

	if got != test.output {
		t.Errorf("AOC2022051() missmatch:\nwant: %s\ngot:  %s", test.output, got)
	}
}
