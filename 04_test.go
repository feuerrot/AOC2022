package main

import "testing"

func TestAOC2022041(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{

		{
			input:  "2-4,6-8",
			output: false,
		},
		{
			input:  "2-3,4-5",
			output: false,
		},
		{
			input:  "5-7,7-9",
			output: false,
		},
		{
			input:  "2-8,3-7",
			output: true,
		},
		{
			input:  "6-6,4-6",
			output: true,
		},
		{
			input:  "2-6,4-8",
			output: false,
		},
	}

	for _, test := range tests {
		got, err := AOC202204Contains(test.input)
		if err != nil {
			t.Errorf("AOC202204Contains(%s) err: %v", test.input, err)
		}
		if got != test.output {
			t.Errorf("AOC202204Contains(%s) missmatch:\nwant: %t\ngot:  %t", test.input, test.output, got)
		}
	}
}
