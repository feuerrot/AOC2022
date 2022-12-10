package main

import "testing"

func TestAOC202210ShortExample(t *testing.T) {
	test := struct {
		input string
		state []struct {
			cycle int
			x     int
		}
	}{
		input: `noop
addx 3
addx -5`,
		state: []struct {
			cycle int
			x     int
		}{
			{cycle: 1, x: 1},
			{cycle: 2, x: 1},
			{cycle: 3, x: 4},
			{cycle: 4, x: 4},
			{cycle: 5, x: -1},
		},
	}

	inst, err := AOC202210ParseInstructions(test.input)
	if err != nil {
		t.Fatalf("AOC202210ParseInstructions(): err: %v", err)
	}

	cpu := AOC202210NewCPU(inst)

	for _, c := range test.state {
		_, err := cpu.Tick()
		if err != nil {
			t.Errorf("AOC202210CPU.Tick() err: %v", err)
		}

		if cpu.Cycle != c.cycle {
			t.Errorf("AOC202210CPU wrong cycle:\nwant: %d\ngot:  %d", c.cycle, cpu.Cycle)
		}

		if cpu.X != c.x {
			t.Errorf("AOC202210CPU wrong register:\nwant: %d\ngot:  %d", c.x, cpu.X)
		}
	}

}

func TestAOC202210LongExample(t *testing.T) {
	test := struct {
		input  string
		output int
		fb     []string
	}{
		input: `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`,
		output: 13140,
		fb: []string{
			"##..##..##..##..##..##..##..##..##..##..",
			"###...###...###...###...###...###...###.",
			"####....####....####....####....####....",
			"#####.....#####.....#####.....#####.....",
			"######......######......######......####",
			"#######.......#######.......#######.....",
		},
	}

	got, err := AOC2022101Helper(test.input)
	if err != nil {
		t.Fatalf("AOC2022101Helper() err: %v", err)
	}

	if got != test.output {
		t.Errorf("AOC2022101Helper() missmatch:\nwant: %d\ngot:  %d", test.output, got)
	}

	gotFb, err := AOC2022102Helper(test.input)
	if err != nil {
		t.Fatalf("AOC2022102Helper() err: %v", err)
	}
	for i := 0; i < len(gotFb); i++ {
		if gotFb[i] != test.fb[i] {
			t.Errorf("AOC2022102Helper() missmatch in line %d:\nwant: %s\ngot:  %s", i, test.fb[i], gotFb[i])
		}
	}

}
