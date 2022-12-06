package main

import "testing"

func TestAOC2022061(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			output: 5,
		},
		{
			input:  "nppdvjthqldpwncqszvftbrmjlhg",
			output: 6,
		},
		{
			input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			output: 10,
		},
		{
			input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			output: 11,
		},
	}

	for _, test := range tests {
		got := AOC202206FindPreamble(test.input)
		if got != test.output {
			t.Errorf("AOC202206FindPreamble(%s) differs:\nwant: %d\ngot:  %d", test.input, test.output, got)
		}
	}
}
