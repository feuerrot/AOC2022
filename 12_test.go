package main

import "testing"

func TestAOC202212ParseMap(t *testing.T) {
	test := struct {
		input    string
		posStart []int
		posEnd   []int
	}{
		input: `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`,
		posStart: []int{0, 0},
		posEnd:   []int{5, 2},
	}

	parsedMap, err := AOC202212ParseMap(test.input)
	if err != nil {
		t.Fatalf("AOC202212ParseMap() err: %v", err)
	}

	nS := parsedMap.NodeStart
	if nS.PosX != test.posStart[0] || nS.PosY != test.posStart[1] {
		t.Errorf("mapStart: %d/%d, testStart: %+v", parsedMap.NodeStart.PosX, parsedMap.NodeStart.PosY, test.posStart)
	}

	nE := parsedMap.NodeEnd
	if nE.PosX != test.posEnd[0] || nE.PosY != test.posEnd[1] {
		t.Errorf("mapEnd: %d/%d, testEnd: %+v", parsedMap.NodeEnd.PosX, parsedMap.NodeEnd.PosY, test.posEnd)
	}

	//if parsedMap.Nodes[2][0].NodeRight != nil {
	//	t.Errorf("Node[2][0] unexpected right neighbor: %+v", parsedMap.Nodes[0][1].NodeRight)
	//}

	//if parsedMap.Nodes[0][0].NodeDown != parsedMap.Nodes[1][0] {
	//	t.Errorf("Node[0][0] unexpected missing downer neighbor: %+v", parsedMap.Nodes[0][0])
	//}

	//way, len := parsedMap.FindWay(parsedMap.NodeStart, []*AOC202212Node{})
	//if len != 31 {
	//	t.Errorf("parsedMap.FindWay() unexpected len %d: Way: %s", len, way)
	//}

	parsedMap.ParseDistances()
	len := parsedMap.SolvePart1()
	if len != 31 {
		t.Errorf("parsedMap.SolvePart1() differs:\nwant: %d\ngot:  %d", 31, len)
	}
}
