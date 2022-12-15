package main

import "testing"

func TestAOC202215(t *testing.T) {
	test := struct {
		input  string
		first  int
		second int
	}{
		input: `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
		first:  26,
		second: 56000011,
	}

	got, err := AOC2022151Helper(test.input, 10)
	if err != nil {
		t.Fatalf("AOC2022151Helper() err: %v", err)
	}

	if got != test.first {
		t.Errorf("AOC2022151Helper() missmatch:\nwant: %d\ngot:  %d", test.first, got)
	}

	got, err = AOC2022152Helper(test.input, 0, 20)
	if err != nil {
		t.Fatalf("AOC2022152Helper() err: %v", err)
	}

	if got != test.second {
		t.Errorf("AOC2022152Helper() missmatch:\nwant: %d\ngot:  %d", test.second, got)
	}
}
