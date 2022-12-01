package main

import "testing"

func TestAOC202201(t *testing.T) {
	var tests = []struct {
		input  string
		output []int
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
			output: []int{24000, 45000},
		},
	}

	for _, test := range tests {
		want := test.output[0]
		got, err := AOC2022011(test.input)
		if err != nil {
			t.Fatalf("AOC2022011() err: %v", err)
		}

		if want != got {
			t.Errorf("AOC2022011() missmatch:\nwant: %d\ngot:  %d", want, got)
		}
	}

	for _, test := range tests {
		want := test.output[1]
		got, err := AOC2022012(test.input)
		if err != nil {
			t.Fatalf("AOC2022012() err: %v", err)
		}

		if want != got {
			t.Errorf("AOC2022012() missmatch:\nwant: %d\ngot:  %d", want, got)
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
			t.Errorf("AOC202201SumMaxN(%+v, %d) err: %v", test.input, test.inputN, err)
		}
		if got != test.output {
			t.Errorf("AOC202201SumMaxN(%+v, %d) missmatch:\nwant: %d\ngot:  %d", test.input, test.inputN, test.output, got)
		}
	}

}
