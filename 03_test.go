package main

import "testing"

func TestAOC2022031(t *testing.T) {
	var tests = []struct {
		input    string
		item     rune
		priority int
	}{
		{
			input:    "vJrwpWtwJgWrhcsFMMfFFhFp",
			item:     'p',
			priority: 16,
		},
		{
			input:    "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			item:     'L',
			priority: 38,
		},
		{
			input:    "PmmdzqPrVvPwwTWBwg",
			item:     'P',
			priority: 42,
		},
		{
			input:    "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			item:     'v',
			priority: 22,
		},
		{
			input:    "ttgJtRGJQctTZtZT",
			item:     't',
			priority: 20,
		},
		{
			input:    "CrZsJsPPZsGzwwsLwLmpwMDw",
			item:     's',
			priority: 19,
		},
	}

	for _, test := range tests {
		got := AOC202203Item(test.input)
		if got != test.item {
			t.Errorf("AOC202203Round(%s) missmatch:\nwant: %v\ngot:  %v", test.input, test.item, got)
		}
		prio := AOC202203Priority(test.item)
		if prio != test.priority {
			t.Errorf("AOC202203Priority(%v) missmatch:\nwant: %d\ngot:  %d", test.item, test.priority, prio)
		}
	}
}

func TestAOC2022032(t *testing.T) {
	var tests = []struct {
		input  []string
		output rune
	}{
		{

			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			output: 'r',
		},
		{
			input: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			output: 'Z',
		},
	}

	for _, test := range tests {
		out := AOC202203Multiitem(test.input[0], test.input[1], test.input[2])
		if out != test.output {
			t.Errorf("AOC202203Item3() missmatch:\nwant: %v\ngot:  %v", test.output, out)
		}
	}
}
