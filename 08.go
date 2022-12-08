package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AOC202208Tree struct {
	Height  int
	Hidden  bool
	HLeft   int
	HRight  int
	HTop    int
	HBottom int
}

func (tree *AOC202208Tree) IsHidden() bool {
	if tree.Height > tree.HLeft {
		return false
	}
	if tree.Height > tree.HRight {
		return false
	}
	if tree.Height > tree.HTop {
		return false
	}
	if tree.Height > tree.HBottom {
		return false
	}

	return true
}

type AOC202208Forest struct {
	Trees [][]*AOC202208Tree
	Rows  int
	Cols  int
}

func AOC202208ParseForest(input string) (*AOC202208Forest, error) {
	forest := AOC202208Forest{}

	for row, line := range strings.Split(input, "\n") {
		if row > forest.Rows {
			forest.Rows = row
		}
		forest.Trees = append(forest.Trees, []*AOC202208Tree{})
		for col, char := range line {
			if col > forest.Cols {
				forest.Cols = col
			}
			height, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("can't parse %s as height: %v", string(char), err)
			}
			forest.Trees[row] = append(forest.Trees[row], &AOC202208Tree{
				Height: height,
			})
		}
	}

	return &forest, nil
}

func (forest *AOC202208Forest) String() string {
	out := []string{""}
	for _, row := range forest.Trees {
		rowOut := []string{}
		for _, tree := range row {
			if tree.Hidden {
				rowOut = append(rowOut, " ")
			} else {
				rowOut = append(rowOut, fmt.Sprintf("%d", tree.Height))
			}
		}
		out = append(out, strings.Join(rowOut, ""))
	}

	return strings.Join(out, "\n")
}

func (forest *AOC202208Forest) HideTrees() {
	topMax := make([]int, forest.Cols)
	for row := 0; row < forest.Rows; row++ {
		leftMax := 0
		for col := 0; col < forest.Cols; col++ {
			if forest.Trees[row][col].HLeft <= leftMax {
				forest.Trees[row][col].HLeft = leftMax
			}
			if forest.Trees[row][col].HTop <= topMax[col] {
				forest.Trees[row][col].HTop = topMax[col]
			}
			if leftMax < forest.Trees[row][col].Height {
				leftMax = forest.Trees[row][col].Height
			}
			if topMax[col] < forest.Trees[row][col].Height {
				topMax[col] = forest.Trees[row][col].Height
			}
		}
	}

	bottomMax := make([]int, forest.Cols)
	for row := forest.Rows; row > 0; row-- {
		rightMax := 0
		for col := forest.Cols; col > 0; col-- {
			if forest.Trees[row][col].HRight <= rightMax {
				forest.Trees[row][col].HRight = rightMax
			}
			if forest.Trees[row][col].HBottom <= bottomMax[col-1] {
				forest.Trees[row][col].HBottom = bottomMax[col-1]
			}

			if row != forest.Rows && row != 0 &&
				col != forest.Cols && col != 0 {
				forest.Trees[row][col].Hidden = forest.Trees[row][col].IsHidden()
			}

			if rightMax < forest.Trees[row][col].Height {
				rightMax = forest.Trees[row][col].Height
			}
			if bottomMax[col-1] < forest.Trees[row][col].Height {
				bottomMax[col-1] = forest.Trees[row][col].Height
			}
		}
	}
}

func (forest *AOC202208Forest) VisibleTrees() int {
	sum := 0
	for _, row := range forest.Trees {
		for _, tree := range row {
			if !tree.Hidden {
				sum += 1
			}
		}
	}

	return sum
}

func AOC2022081(input string) (string, error) {
	forest, err := AOC202208ParseForest(input)
	if err != nil {
		return "", err
	}
	forest.HideTrees()

	return fmt.Sprintf("%d", forest.VisibleTrees()), nil
}
