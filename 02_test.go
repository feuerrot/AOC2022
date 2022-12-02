package main

import "testing"

func TestAOC202202(t *testing.T) {
	input := `
A Y
B X
C Z
`
	want := 15
	got, err := AOC2022021(input)
	if err != nil {
		t.Fatalf("AOC2022021() err: %v", err)
	}

	if want != got {
		t.Errorf("AOC2022021() missmatch:\nwant: %d\ngot:  %d", want, got)
	}
}
