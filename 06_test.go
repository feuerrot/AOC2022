package main

import (
	"os"
	"testing"
)

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
			t.Errorf("AOC202206FindUnique(%s, 4) differs:\nwant: %d\ngot:  %d", test.input, test.preamble, got)
		}

		got = AOC202206FindUniqueJump(test.input, 4)
		if got != test.preamble {
			t.Errorf("AOC202206FindUniqueJump(%s, 4) differs:\nwant: %d\ngot:  %d", test.input, test.preamble, got)
		}

		got = AOC202206FindUniqueReverseJump(test.input, 4)
		if got != test.preamble {
			t.Errorf("AOC202206FindUniqueReverseJump(%s, 4) differs:\nwant: %d\ngot:  %d", test.input, test.preamble, got)
		}

		got = AOC202206FindUnique(test.input, 14)
		if got != test.message {
			t.Errorf("AOC202206FindUnique(%s, 14) differs:\nwant: %d\ngot:  %d", test.input, test.message, got)
		}

		got = AOC202206FindUniqueJump(test.input, 14)
		if got != test.message {
			t.Errorf("AOC202206FindUniqueJump(%s, 14) differs:\nwant: %d\ngot:  %d", test.input, test.message, got)
		}

		got = AOC202206FindUniqueReverseJump(test.input, 14)
		if got != test.message {
			t.Errorf("AOC202206FindUniqueReverseJump(%s, 14) differs:\nwant: %d\ngot:  %d", test.input, test.message, got)
		}
	}
}

func BenchmarkAOC202206FindUnique(b *testing.B) {
	data, err := os.ReadFile("input/202206")
	if err != nil {
		b.Fatalf("can't open input: %v", err)
	}
	input := string(data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AOC202206FindUnique(input, 14)
	}
}

func BenchmarkAOC202206FindUniqueJump(b *testing.B) {
	data, err := os.ReadFile("input/202206")
	if err != nil {
		b.Fatalf("can't open input: %v", err)
	}
	input := string(data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AOC202206FindUniqueJump(input, 14)
	}
}

func BenchmarkAOC202206FindUniqueReverseJump(b *testing.B) {
	data, err := os.ReadFile("input/202206")
	if err != nil {
		b.Fatalf("can't open input: %v", err)
	}
	input := string(data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AOC202206FindUniqueReverseJump(input, 14)
	}
}
