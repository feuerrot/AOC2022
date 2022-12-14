package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type AOC202212Node struct {
	Height    rune
	PosX      int
	PosY      int
	NodeUp    *AOC202212Node
	NodeDown  *AOC202212Node
	NodeLeft  *AOC202212Node
	NodeRight *AOC202212Node
}

type AOC202212Way []*AOC202212Node

func (way AOC202212Way) String() string {
	parts := []string{}
	for _, node := range way {
		parts = append(parts, fmt.Sprintf("N %d/%d H %d", node.PosX, node.PosY, node.Height))
	}

	return strings.Join(parts, " -> ")
}

func (n *AOC202212Node) NeighborReachable(neighbor *AOC202212Node) bool {
	if neighbor == nil {
		return false
	}

	selfHeight := n.Height
	if selfHeight == 'S' {
		selfHeight = 'a'
	} else if selfHeight == 'E' {
		selfHeight = 'z'
	}
	neighborHeight := neighbor.Height
	if neighborHeight == 'S' {
		neighborHeight = 'a'
	} else if neighborHeight == 'E' {
		neighborHeight = 'z'
	}

	return neighborHeight <= selfHeight+1
}

type AOC202212Map struct {
	NodeStart *AOC202212Node
	NodeEnd   *AOC202212Node
	SizeX     int
	SizeY     int
	Nodes     [][]*AOC202212Node
	Distance  map[int][]*AOC202212Node
}

func (m *AOC202212Map) NodeLeft(x, y int) *AOC202212Node {
	if x-1 < 0 {
		return nil
	}
	return m.Nodes[y][x-1]
}

func (m *AOC202212Map) NodeRight(x, y int) *AOC202212Node {
	if x >= m.SizeX {
		return nil
	}
	return m.Nodes[y][x+1]
}

func (m *AOC202212Map) NodeUp(x, y int) *AOC202212Node {
	if y-1 < 0 {
		return nil
	}
	return m.Nodes[y-1][x]
}

func (m *AOC202212Map) NodeDown(x, y int) *AOC202212Node {
	if y >= m.SizeY {
		return nil
	}
	return m.Nodes[y+1][x]
}

func (m *AOC202212Map) FindWay(current *AOC202212Node, previous AOC202212Way) (AOC202212Way, int) {
	search := []*AOC202212Node{
		current.NodeUp,
		current.NodeDown,
		current.NodeLeft,
		current.NodeRight,
	}

	sort.Slice(search, func(i, j int) bool {
		if search[i] == nil || search[j] == nil {
			return false
		}
		if search[i].PosX > search[j].PosX {
			return true
		}
		if search[i].PosY > search[j].PosY {
			return true
		}

		if search[i].Height != current.Height {
			return true
		}
		return false
	})

	bestLen := -1
	var bestChoice []*AOC202212Node
	var wg sync.WaitGroup
	var bcm sync.Mutex

CHOICE:
	for _, choice := range search {
		if choice == nil {
			continue
		}

		if choice == m.NodeEnd {
			return []*AOC202212Node{choice}, 1
		}

		for _, prev := range previous {
			if choice == prev {
				continue CHOICE
			}
		}

		lchoice := choice
		previous := previous
		wg.Add(1)
		go func() {
			previous = append(previous, current)
			nextWay, nextLen := m.FindWay(lchoice, previous)
			bcm.Lock()
			if nextWay != nil && (bestLen == -1 || nextLen < bestLen) {
				bestChoice = nextWay
				bestLen = nextLen
			}
			bcm.Unlock()
			wg.Done()
		}()
		if len(previous) > 2 {
			wg.Wait()
		}
	}

	wg.Wait()

	if bestChoice == nil {
		return nil, 0
	}

	bestChoice = append(bestChoice, current)

	return bestChoice, bestLen + 1
}

func (m *AOC202212Map) ParseDistances() {
	found := []*AOC202212Node{m.NodeStart}
	distance := map[int][]*AOC202212Node{
		0: {m.NodeStart},
	}

	for i := 1; ; i++ {
		curDist := []*AOC202212Node{}
		for _, node := range distance[i-1] {
			if node == nil {
				continue
			}
			search := []*AOC202212Node{
				node.NodeUp,
				node.NodeDown,
				node.NodeLeft,
				node.NodeRight,
			}

		NEIGHBOR:
			for _, neighbor := range search {
				for _, prevFound := range found {
					if neighbor == prevFound {
						continue NEIGHBOR
					}
				}
				curDist = append(curDist, neighbor)
				found = append(found, neighbor)
			}
		}
		if len(curDist) == 0 {
			break
		}
		distance[i] = curDist
	}
	m.Distance = distance
}

func (m *AOC202212Map) SolvePart1() int {
	distance := []int{}
	for i := range m.Distance {
		distance = append(distance, i)
	}

	sort.Ints(distance)
	for _, i := range distance {
		for _, node := range m.Distance[i] {
			if node == m.NodeEnd {
				return i
			}
		}
	}

	return 0
}

func AOC202212ParseMap(input string) (AOC202212Map, error) {
	rtn := AOC202212Map{
		Nodes: make([][]*AOC202212Node, 0),
	}

	for y, line := range strings.Split(input, "\n") {
		if y > rtn.SizeY {
			rtn.SizeY = y
		}
		lineNodes := make([]*AOC202212Node, 0)
		for x, char := range line {
			if x > rtn.SizeX {
				rtn.SizeX = x
			}
			node := &AOC202212Node{
				Height: char,
				PosX:   x,
				PosY:   y,
			}
			if char == 'S' {
				rtn.NodeStart = node
				//node.Height = 'a'
			} else if char == 'E' {
				rtn.NodeEnd = node
				//node.Height = 'z'
			}
			lineNodes = append(lineNodes, node)
		}
		rtn.Nodes = append(rtn.Nodes, lineNodes)
	}

	for y, row := range rtn.Nodes {
		for x, node := range row {
			if node.NeighborReachable(rtn.NodeLeft(x, y)) {
				rtn.Nodes[y][x].NodeLeft = rtn.NodeLeft(x, y)
			}
			if node.NeighborReachable(rtn.NodeRight(x, y)) {
				rtn.Nodes[y][x].NodeRight = rtn.NodeRight(x, y)
			}
			if node.NeighborReachable(rtn.NodeUp(x, y)) {
				rtn.Nodes[y][x].NodeUp = rtn.NodeUp(x, y)
			}
			if node.NeighborReachable(rtn.NodeDown(x, y)) {
				rtn.Nodes[y][x].NodeDown = rtn.NodeDown(x, y)
			}
		}
	}

	return rtn, nil
}

func AOC2022121(input string) (string, error) {
	parsedMap, err := AOC202212ParseMap(input)
	if err != nil {
		return "", err
	}
	parsedMap.ParseDistances()
	len := parsedMap.SolvePart1()
	return fmt.Sprintf("%d", len), nil
}
