package main

import "testing"

func TestAOC2022011(t *testing.T) {
	input := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	want := 24000
	got, err := AOC2022011(input)
	if err != nil {
		t.Fatalf("AOC2022011() err: %v", err)
	}

	if want != got {
		t.Errorf("AOC2022011() missmatch:\nwant: %d\ngot:  %d", want, got)
	}

}
