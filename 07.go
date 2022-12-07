package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type AOC202207Node struct {
	Name   string
	Size   int
	Child  []*AOC202207Node
	Parent *AOC202207Node
}

func (node *AOC202207Node) Path() string {
	if node.Name == "" {
		return ""
	}
	return node.Parent.Path() + "/" + node.Name
}

func (node *AOC202207Node) String() string {
	rtn := []string{
		fmt.Sprintf("%s: %d", node.Path(), node.GetSize()),
	}

	for _, child := range node.Child {
		rtn = append(rtn, child.String())
	}

	return strings.Join(rtn, "\n")
}

func (node *AOC202207Node) FindChild(name string) (*AOC202207Node, error) {
	for _, child := range node.Child {
		if child.Name == name {
			return child, nil
		}
	}
	return nil, fmt.Errorf("node %s has no child %s", node.Name, name)
}

func (node *AOC202207Node) CreateDirectory(name string) {
	node.Child = append(node.Child, &AOC202207Node{
		Name:   name,
		Child:  []*AOC202207Node{},
		Parent: node,
	})
}

func (node *AOC202207Node) CreateFile(name string, size int) {
	node.Child = append(node.Child, &AOC202207Node{
		Name:   name,
		Parent: node,
		Size:   size,
	})
}

func (node *AOC202207Node) GetSize() int {
	size := 0
	for _, child := range node.Child {
		if child.Child == nil {
			size += child.Size
		} else {
			size += child.GetSize()
		}
	}

	node.Size = size

	return size
}

func (node *AOC202207Node) GetDirsizes() []int {
	rtn := []int{}
	rtn = append(rtn, node.Size)
	for _, child := range node.Child {
		if child.Child != nil {
			rtn = append(rtn, child.GetDirsizes()...)
		}
	}

	return rtn
}

type AOC202207FS struct {
	Root     *AOC202207Node
	Position *AOC202207Node
}

func (fs *AOC202207FS) String() string {
	return fs.Root.String()
}

func (fs *AOC202207FS) Size() int {
	return fs.Root.GetSize()
}

func (fs *AOC202207FS) GetNode(path string) (*AOC202207Node, error) {
	parts := strings.Split(path, "/")
	if parts[0] != "" {
		return nil, fmt.Errorf("can't find node without a full path")
	}

	curNode := fs.Root
	for _, part := range parts[1:] {
		next, err := curNode.FindChild(part)
		if err != nil {
			return nil, fmt.Errorf("%s has no child %s", curNode.Path(), part)
		}
		curNode = next
	}

	return curNode, nil
}

func (fs *AOC202207FS) GetDirsizes() []int {
	return fs.Root.GetDirsizes()
}

func (fs *AOC202207FS) Populate(input string) error {
	dirmode := false
	for _, line := range strings.Split(input, "\n") {
		if line == "$ ls" {
			dirmode = true
			continue
		}
		if line == "$ cd /" {
			dirmode = false
			fs.Position = fs.Root
			continue
		}
		if line == "$ cd .." {
			dirmode = false
			fs.Position = fs.Position.Parent
			continue
		}
		if line[:5] == "$ cd " {
			dirmode = false
			target := strings.Split(line, " ")[2]
			child, err := fs.Position.FindChild(target)
			if err != nil {
				return fmt.Errorf("node %s: can't switch to child %s: %v", fs.Position.Name, target, err)
			}
			fs.Position = child
			continue
		}

		if dirmode {
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				fs.Position.CreateDirectory(parts[1])
			} else {
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					return fmt.Errorf("can't convert size of \"%s\" into an integer: %v", line, err)
				}
				fs.Position.CreateFile(parts[1], size)
			}
			continue
		}

		return fmt.Errorf("unexpected entry: %s", line)
	}

	fs.Size()
	return nil
}

func AOC202207NewFS() (*AOC202207FS, error) {
	fs := &AOC202207FS{
		Root: &AOC202207Node{
			Name:  "",
			Child: []*AOC202207Node{},
		},
	}

	fs.Position = fs.Root

	return fs, nil
}

func (fs *AOC202207FS) AOC2022071Helper() int {
	sum := 0
	for _, elem := range fs.GetDirsizes() {
		if elem <= 100000 {
			sum += elem
		}
	}
	return sum
}

func AOC202207FSInit(input string) (*AOC202207FS, error) {
	fs, err := AOC202207NewFS()
	if err != nil {
		return nil, fmt.Errorf("can't init new FS: %v", err)
	}
	fs.Populate(input)

	return fs, nil
}

func AOC2022071(input string) (string, error) {
	fs, err := AOC202207FSInit(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", fs.AOC2022071Helper()), nil
}

func (fs *AOC202207FS) AOC2022072Helper() int {
	space := 70000000
	requiredFree := 30000000
	currentFree := space - fs.Size()
	delta := requiredFree - currentFree

	dirsizes := fs.GetDirsizes()
	sort.Ints(dirsizes)
	for _, size := range dirsizes {
		if size < delta {
			continue
		}
		return size
	}

	return 0
}

func AOC2022072(input string) (string, error) {
	fs, err := AOC202207FSInit(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", fs.AOC2022072Helper()), nil
}
