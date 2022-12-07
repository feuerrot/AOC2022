package main

import "testing"

func TestAOC202207ParseTree(t *testing.T) {
	test := struct {
		input  string
		output int
	}{
		input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
		output: 48381165,
	}

	fs, err := AOC202207NewFS()
	if err != nil {
		t.Fatalf("can't create FS: %v", err)
	}

	err = fs.Populate(test.input)
	if err != nil {
		t.Fatalf("can't populate FS: %v", err)
	}

	got := fs.Size()
	if got != test.output {
		t.Errorf("AOC202207FS.Size() missmatch:\nwant: %d\ngot:  %d", test.output, got)
	}

	dirE, err := fs.GetNode("/a/e")
	if err != nil {
		t.Fatalf("can't access directory /a/e: %v", err)
	}
	dirESize := dirE.GetSize()
	if dirESize != 584 {
		t.Errorf("dirE.GetSize() missmatch:\nwant: 584\ngot:  %d", dirESize)
	}

	dirA, err := fs.GetNode("/a")
	if err != nil {
		t.Fatalf("can't access directory /a: %v", err)
	}
	dirASize := dirA.GetSize()
	if dirASize != 94853 {
		t.Errorf("dirA.GetSize() missmatch:\nwant: 94853\ngot:  %d", dirASize)
	}

	part1 := fs.AOC2022071Helper()
	if part1 != 95437 {
		t.Errorf("AOC2022071Helper() missmatch:\nwant: 95437\ngot:  %d", part1)
	}
}
