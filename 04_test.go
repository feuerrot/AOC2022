package main

import "testing"

func TestAOC2022041(t *testing.T) {
	var tests = []struct {
		input          string
		fullOverlap    bool
		partialOverlap bool
	}{

		{
			input:          "2-4,6-8",
			fullOverlap:    false,
			partialOverlap: false,
		},
		{
			input:          "2-3,4-5",
			fullOverlap:    false,
			partialOverlap: false,
		},
		{
			input:          "5-7,7-9",
			fullOverlap:    false,
			partialOverlap: true,
		},
		{
			input:          "2-8,3-7",
			fullOverlap:    true,
			partialOverlap: true,
		},
		{
			input:          "6-6,4-6",
			fullOverlap:    true,
			partialOverlap: true,
		},
		{
			input:          "2-6,4-8",
			fullOverlap:    false,
			partialOverlap: true,
		},
	}

	for _, test := range tests {
		left, right, err := AOC202204Parse(test.input)
		if err != nil {
			t.Errorf("AOC202204Parse(%s) err: %v", test.input, err)
		}

		got := AOC202204FullOverlap(left, right)
		if got != test.fullOverlap {
			t.Errorf("AOC202204FullOverlap(%s) missmatch:\nwant: %t\ngot:  %t", test.input, test.fullOverlap, got)
		}

		got = AOC202204PartialOverlap(left, right)
		if got != test.partialOverlap {
			t.Errorf("AOC202204PartialOverlap(%s) missmatch:\nwant: %t\ngot:  %t", test.input, test.partialOverlap, got)
		}
	}
}
