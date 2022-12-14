package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MapPoint int

const (
	Air MapPoint = iota
	Rock
	Sand
)

type AOC202214SandMap struct {
	// Point[y][x]
	Point map[int]map[int]MapPoint
	MaxY  int
	MinX  int
	MaxX  int
}

func (sm *AOC202214SandMap) DrawMap() string {
	rtn := []string{}
	for y := 0; y <= sm.MaxY; y++ {
		line := ""
		for x := sm.MinX; x <= sm.MaxX; x++ {
			if x == 500 && y == 0 {
				line += "+"
				continue
			}
			pt := sm.GetPoint(x, y)
			if pt == Air {
				line += "."
			} else if pt == Rock {
				line += "#"
			} else if pt == Sand {
				line += "o"
			}
		}
		rtn = append(rtn, line)
	}

	return strings.Join(rtn, "\n")
}

func (sm *AOC202214SandMap) SetPoint(x, y int, pt MapPoint) {
	if sm.Point[y] == nil {
		sm.Point[y] = make(map[int]MapPoint, 0)
	}

	sm.Point[y][x] = pt
}

func (sm *AOC202214SandMap) GetPoint(x, y int) MapPoint {
	if sm.Point[y] == nil {
		sm.Point[y] = make(map[int]MapPoint)
		return Air
	}

	elem, ok := sm.Point[y][x]
	if !ok {
		return Air
	}
	return elem
}

func (sm *AOC202214SandMap) AddSand() bool {
	xPos, yPos := 500, 0
	for {
		// check if it fell off
		if yPos > sm.MaxY {
			return false
		}

		// check below
		if sm.GetPoint(xPos, yPos+1) == Air {
			yPos += 1
			continue
		}

		// check below & left
		if sm.GetPoint(xPos-1, yPos+1) == Air {
			xPos -= 1
			yPos += 1
			continue
		}

		if sm.GetPoint(xPos+1, yPos+1) == Air {
			xPos += 1
			yPos += 1
			continue
		}

		sm.SetPoint(xPos, yPos, Sand)
		return true
	}
}

func (sm *AOC202214SandMap) ParseRocks(input string) error {
	for _, line := range strings.Split(input, "\n") {
		points := [][]int{}
		for _, part := range strings.Split(line, " -> ") {
			coord := strings.Split(part, ",")
			coordX, err := strconv.Atoi(coord[0])
			if err != nil {
				return fmt.Errorf("can't parse coordinate %s: %v", part, err)
			}
			coordY, err := strconv.Atoi(coord[1])
			if err != nil {
				return fmt.Errorf("can't parse coordinate %s: %v", part, err)
			}
			if coordY > sm.MaxY {
				sm.MaxY = coordY
			}
			if coordX < sm.MinX || sm.MinX == 0 {
				sm.MinX = coordX
			}
			if coordX > sm.MaxX {
				sm.MaxX = coordX
			}

			points = append(points, []int{coordX, coordY})
		}

		for i := 0; i < len(points)-1; i++ {
			startX, startY := points[i][0], points[i][1]
			endX, endY := points[i+1][0], points[i+1][1]

			if startX == endX {
				if startY > endY {
					startY, endY = endY, startY
				}
				for y := startY; y <= endY; y++ {
					sm.SetPoint(startX, y, Rock)
				}
			} else {
				if startX > endX {
					startX, endX = endX, startX
				}
				for x := startX; x <= endX; x++ {
					sm.SetPoint(x, startY, Rock)
				}
			}
		}
	}

	return nil
}

func AOC2022141Helper(input string) (int, error) {
	sandMap := *&AOC202214SandMap{
		Point: make(map[int]map[int]MapPoint),
	}

	err := sandMap.ParseRocks(input)
	if err != nil {
		return 0, fmt.Errorf("can't parse rocks: %v", err)
	}

	log.Printf("Map before:\n%s", sandMap.DrawMap())

	grains := 0
	for {
		if !sandMap.AddSand() {
			break
		}
		grains += 1
	}

	log.Printf("Map after:\n%s", sandMap.DrawMap())

	return grains, nil
}

func AOC2022141(input string) (string, error) {
	grains, err := AOC2022141Helper(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", grains), nil
}
