package main

import "testing"

func TestAOC2022061(t *testing.T) {
	tests := []struct {
		input    string
		preamble int
		message  int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			preamble: 7,
			message:  19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			preamble: 5,
			message:  23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			preamble: 6,
			message:  23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			preamble: 10,
			message:  29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			preamble: 11,
			message:  26,
		},
	}

	for _, test := range tests {
		got := AOC202206FindUnique(test.input, 4)
		if got != test.preamble {
			t.Errorf("AOC202206FindPreamble(%s) differs:\nwant: %d\ngot:  %d", test.input, test.preamble, got)
		}

		got = AOC202206FindUnique(test.input, 14)
		if got != test.message {
			t.Errorf("AOC202206FindMessage(%s) differs:\nwant: %d\ngot:  %d", test.input, test.message, got)
		}
	}
}
