package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAOC202201Helper(t *testing.T) {
	var tests = []struct {
		input     string
		output    []int
		outputErr bool
	}{
		{
			input: `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			output: []int{4000, 6000, 10000, 11000, 24000},
		},
		{
			input:     ``,
			output:    nil,
			outputErr: false,
		},
		{
			input:     `Pinselohrkatze`,
			output:    nil,
			outputErr: true,
		},
	}

	for _, test := range tests {
		want := test.output
		got, err := AOC202201Helper(test.input)
		if err != nil && !test.outputErr {
			t.Fatalf("AOC2022011Helper(%+v) err: %v", test.input, err)
		}
		if err == nil && test.outputErr {
			t.Fatalf("AOC2022011Helper(%+v) should return an error instead of %d", test.input, test.output)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("AOC2022011Helper(%+v) missmatch (-want +got):\n%s", test.input, diff)
		}
	}
}

func TestAOC202201SumMaxN(t *testing.T) {
	var tests = []struct {
		input     []int
		inputN    int
		output    int
		outputErr bool
	}{
		{
			input:     []int{1, 2, 3, 4},
			inputN:    1,
			output:    4,
			outputErr: false,
		},
		{
			input:     []int{1, 2, 3, 4},
			inputN:    4,
			output:    10,
			outputErr: false,
		},
		{
			input:     []int{1, 2, 3, 4},
			inputN:    0,
			output:    0,
			outputErr: false,
		},
		{
			input:     []int{1, 2, 3, 4},
			inputN:    5,
			output:    0,
			outputErr: true,
		},
	}

	for _, test := range tests {
		got, err := AOC202201SumMaxN(test.input, test.inputN)
		if err != nil && !test.outputErr {
			t.Fatalf("AOC202201SumMaxN(%+v, %d) err: %v", test.input, test.inputN, err)
		}
		if err == nil && test.outputErr {
			t.Fatalf("AOC202201SumMaxN(%+v, %d) should return an error instead of %d", test.input, test.inputN, test.output)
		}

		if got != test.output {
			t.Errorf("AOC202201SumMaxN(%+v, %d) missmatch:\nwant: %d\ngot:  %d", test.input, test.inputN, test.output, got)
		}
	}

}
