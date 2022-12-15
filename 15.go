package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type AOC202215MapPoint uint8
type AOC202215MapPointUsed bool
type AOC202215MapPointsUsed []AOC202215MapPointUsed

const (
	MPEmpty AOC202215MapPoint = iota
	MPBeacon
	MPSensor
	MPNoBeacon
)

type AOC202215SensorBeacon struct {
	Sensor   []int
	Beacon   []int
	Distance int
	Border   [][]int
}

func (sb *AOC202215SensorBeacon) Overlap(sensor *AOC202215SensorBeacon) bool {
	sensorDistance := absInt(sb.Sensor[0]-sensor.Sensor[0]) + absInt(sb.Sensor[1]-sensor.Sensor[1])
	maxDistance := sb.Distance + sensor.Distance
	return sensorDistance <= maxDistance
}

func (sb *AOC202215SensorBeacon) Contains(x, y int) bool {
	return (absInt(sb.Sensor[0]-x)+absInt(sb.Sensor[1]-y) <= sb.Distance)
}

type AOC202215Map struct {
	SensorBeacons []*AOC202215SensorBeacon
	Map           map[int]map[int]AOC202215MapPoint
	IntMap        [][]AOC202215MapPointUsed
	RowMap        AOC202215MapPointsUsed
}

func (m *AOC202215Map) PrintIntMap() string {
	rows := []string{}
	for _, row := range m.IntMap {
		rows = append(rows, fmt.Sprintf("%+v", row))
	}
	return strings.Join(rows, "\n")
}

func (m *AOC202215Map) SetPoint(x, y int, pt AOC202215MapPoint) {
	if m.Map == nil {
		m.Map = make(map[int]map[int]AOC202215MapPoint)
	}
	if _, ok := m.Map[y]; !ok {
		m.Map[y] = make(map[int]AOC202215MapPoint)
	}

	m.Map[y][x] = pt
}

func (m *AOC202215Map) GetRow(y int) map[int]AOC202215MapPoint {
	row, ok := m.Map[y]
	if !ok {
		return nil
	}

	return row
}

func (m *AOC202215Map) PopulateMap() {
	for _, sb := range m.SensorBeacons {
		for y := sb.Sensor[1] - sb.Distance; y <= sb.Sensor[1]+sb.Distance; y++ {
			xDelta := sb.Distance - absInt(sb.Sensor[1]-y)
			for x := sb.Sensor[0] - xDelta; x <= sb.Sensor[0]+xDelta; x++ {
				m.SetPoint(x, y, MPNoBeacon)
			}
		}
		m.SetPoint(sb.Beacon[0], sb.Beacon[1], MPBeacon)
		m.SetPoint(sb.Sensor[0], sb.Sensor[1], MPSensor)
	}
}

func (m *AOC202215Map) PopulateMapROI(row int) {
	for _, sb := range m.SensorBeacons {
		yMin := sb.Sensor[1] - sb.Distance
		yMax := sb.Sensor[1] + sb.Distance
		if yMin > row || yMax < row {
			continue
		}

		xDelta := sb.Distance - absInt(sb.Sensor[1]-row)
		for x := sb.Sensor[0] - xDelta; x <= sb.Sensor[0]+xDelta; x++ {
			m.SetPoint(x, row, MPNoBeacon)
		}

		m.SetPoint(sb.Beacon[0], sb.Beacon[1], MPBeacon)
		m.SetPoint(sb.Sensor[0], sb.Sensor[1], MPSensor)
	}
}

func (m *AOC202215Map) GetMapROI(row int) map[int]AOC202215MapPointUsed {
	rowMap := make(map[int]AOC202215MapPointUsed)
	for _, sb := range m.SensorBeacons {
		yMin := sb.Sensor[1] - sb.Distance
		yMax := sb.Sensor[1] + sb.Distance
		if yMin > row || yMax < row {
			continue
		}

		xDelta := sb.Distance - absInt(sb.Sensor[1]-row)
		for x := sb.Sensor[0] - xDelta; x <= sb.Sensor[0]+xDelta; x++ {
			rowMap[x] = true
		}
		if sb.Beacon[1] == row {
			rowMap[sb.Beacon[0]] = true
		}
		if sb.Sensor[1] == row {
			rowMap[sb.Sensor[0]] = true
		}
	}

	return rowMap
}

func (m *AOC202215Map) PopulateMapSparse(min, max int) {
	for _, sb := range m.SensorBeacons {
		for y := sb.Sensor[1] - sb.Distance; y <= sb.Sensor[1]+sb.Distance; y++ {
			if y < min || y > max {
				continue
			}

			xDelta := sb.Distance - absInt(sb.Sensor[1]-y)
			for _, x := range []int{sb.Sensor[0] - xDelta, sb.Sensor[0] + xDelta} {
				if x < min || x > max {
					continue
				}
				m.SetPoint(x, y, MPNoBeacon)
			}
		}
		m.SetPoint(sb.Beacon[0], sb.Beacon[1], MPBeacon)
		m.SetPoint(sb.Sensor[0], sb.Sensor[1], MPSensor)
	}
}

func (m *AOC202215Map) PopulateIntMap(min, max int) {
	m.IntMap = make([][]AOC202215MapPointUsed, max)
	for _, sb := range m.SensorBeacons {
		for y := sb.Sensor[1] - sb.Distance; y <= sb.Sensor[1]+sb.Distance; y++ {
			if y < min || y >= max {
				continue
			}

			if m.IntMap[y] == nil {
				m.IntMap[y] = make([]AOC202215MapPointUsed, max)
			}

			xDelta := sb.Distance - absInt(sb.Sensor[1]-y)
			for x := sb.Sensor[0] - xDelta; x <= sb.Sensor[0]+xDelta; x++ {
				if x < min || x >= max {
					continue
				}
				m.IntMap[y][x] = true
			}
		}
		m.IntMap[sb.Beacon[1]][sb.Beacon[0]] = true
		m.IntMap[sb.Sensor[1]][sb.Sensor[0]] = true
	}
}

func (m *AOC202215Map) PopulateIntMapROI(row, min, max int) {
	rowMap := make([]AOC202215MapPointUsed, max+2)
	for _, sb := range m.SensorBeacons {
		yMin := sb.Sensor[1] - sb.Distance
		yMax := sb.Sensor[1] + sb.Distance
		if yMin > row || yMax < row {
			continue
		}

		xDelta := sb.Distance - absInt(sb.Sensor[1]-row)
		xMin := sb.Sensor[0] - xDelta
		if xMin < min {
			xMin = min
		}
		xMax := sb.Sensor[0] + xDelta
		if xMax > max {
			xMax = max
		}
		for x := xMin; x <= xMax; x++ {
			rowMap[x] = true
		}

		if sb.Beacon[1] == row && sb.Beacon[0] > min && sb.Beacon[0] < max {
			rowMap[sb.Beacon[0]] = true
		}

		if sb.Sensor[1] == row && sb.Sensor[0] > min && sb.Sensor[0] < max {
			rowMap[sb.Sensor[0]] = true
		}
	}
	m.RowMap = rowMap
}

func (m *AOC202215Map) GetIntMapROI(row, min, max int) AOC202215MapPointsUsed {
	rowMap := make(AOC202215MapPointsUsed, max+1)
	for _, sb := range m.SensorBeacons {
		yMin := sb.Sensor[1] - sb.Distance
		yMax := sb.Sensor[1] + sb.Distance
		if yMin > row || yMax < row {
			continue
		}

		xDelta := sb.Distance - absInt(sb.Sensor[1]-row)
		xMin := sb.Sensor[0] - xDelta
		if xMin < min {
			xMin = min
		}
		xMax := sb.Sensor[0] + xDelta
		if xMax > max {
			xMax = max
		}
		for x := xMin; x <= xMax; x++ {
			rowMap[x] = true
		}

		if sb.Beacon[1] == row && sb.Beacon[0] > min && sb.Beacon[0] < max {
			rowMap[sb.Beacon[0]] = true
		}

		if sb.Sensor[1] == row && sb.Sensor[0] > min && sb.Sensor[0] < max {
			rowMap[sb.Sensor[0]] = true
		}
	}

	return rowMap
}

func (m *AOC202215Map) PopulateBorder(min, max int) {
	for _, sb := range m.SensorBeacons {
		borderDistance := sb.Distance + 1
		for y := sb.Sensor[1] - borderDistance; y <= sb.Sensor[1]+borderDistance; y++ {
			if y < min || y > max {
				continue
			}

			xDelta := borderDistance - absInt(sb.Sensor[1]-y)
			for _, x := range []int{sb.Sensor[0] - xDelta, sb.Sensor[0] + xDelta} {
				if x < min || x > max {
					continue
				}
				sb.Border = append(sb.Border, []int{x, y})
			}
		}
	}
}

func (m *AOC202215Map) ParseInput(input string) error {
	m.SensorBeacons = []*AOC202215SensorBeacon{}
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, line := range strings.Split(input, "\n") {
		parts := re.FindStringSubmatch(line)
		if len(parts) != 5 {
			return fmt.Errorf("can't parse line %s", line)
		}
		intParts := []int{}
		for _, part := range parts[1:] {
			intPart, err := strconv.Atoi(part)
			if err != nil {
				return fmt.Errorf("can't convert %s to int: %v", part, err)
			}
			intParts = append(intParts, intPart)
		}

		sb := &AOC202215SensorBeacon{
			Sensor:   intParts[:2],
			Beacon:   intParts[2:],
			Distance: absInt(intParts[0]-intParts[2]) + absInt(intParts[1]-intParts[3]),
			Border:   [][]int{},
		}
		m.SensorBeacons = append(m.SensorBeacons, sb)
	}

	return nil
}

func AOC2022151Helper(input string, row int) (int, error) {
	sbMap := &AOC202215Map{}
	err := sbMap.ParseInput(input)
	if err != nil {
		return 0, fmt.Errorf("can't parse input: %v", err)
	}
	sbMap.PopulateMapROI(row)

	count := 0
	mapRow := sbMap.GetRow(row)
	if mapRow == nil {
		return 0, fmt.Errorf("row %d is empty", row)
	}
	for k := range mapRow {
		if mapRow[k] != MPBeacon {
			count += 1
		}
	}

	return count, nil
}

func AOC2022151(input string) (string, error) {
	res, err := AOC2022151Helper(input, 2000000)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", res), nil
}

func AOC2022152Helper(input string, min, max int) (int, error) {
	sbMap := &AOC202215Map{}
	err := sbMap.ParseInput(input)
	if err != nil {
		return 0, fmt.Errorf("can't parse input: %v", err)
	}

	mtx := sync.Mutex{}

	checkCollision := func(left, right *AOC202215SensorBeacon, match *[][]int) {
		colmap := map[string]int{}
		for _, ptLeft := range left.Border {
			key := fmt.Sprintf("%d|%d", ptLeft[0], ptLeft[1])
			if _, ok := colmap[key]; !ok {
				colmap[key] = 1
			} else {
				colmap[key] += 1
			}
		}
		for _, ptRight := range right.Border {
			key := fmt.Sprintf("%d|%d", ptRight[0], ptRight[1])
			if _, ok := colmap[key]; ok {
				colmap[key] += 1
			}
		}

	OUTER:
		for key, entry := range colmap {
			if entry == 2 {
				parts := strings.Split(key, "|")
				x, _ := strconv.Atoi(parts[0])
				y, _ := strconv.Atoi(parts[1])
				mtx.Lock()
				for _, m := range *match {
					if m[0] == x && m[1] == y {
						mtx.Unlock()
						continue OUTER
					}
				}
				(*match) = append((*match), []int{x, y})
				mtx.Unlock()
			}
		}
	}

	sbMap.PopulateBorder(min, max)

	lrChan := make(chan []int)
	match := [][]int{}

	gowg := sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		gowg.Add(1)
		go func() {
			for lr := range lrChan {
				checkCollision(sbMap.SensorBeacons[lr[0]], sbMap.SensorBeacons[lr[1]], &match)
			}
			gowg.Done()
		}()
	}

	for left := 0; left < len(sbMap.SensorBeacons)-1; left++ {
		for right := left + 1; right < len(sbMap.SensorBeacons); right++ {
			if !sbMap.SensorBeacons[left].Overlap(sbMap.SensorBeacons[right]) {
				continue
			}
			lrChan <- []int{left, right}
		}
	}
	close(lrChan)

	gowg.Wait()

	filteredMatches := [][]int{}
OUTER:
	for _, m := range match {
		for _, beacon := range sbMap.SensorBeacons {
			if beacon.Contains(m[0], m[1]) {
				continue OUTER
			}
		}
		filteredMatches = append(filteredMatches, m)
	}

	if len(filteredMatches) == 1 {
		return filteredMatches[0][0]*4000000 + filteredMatches[0][1], nil
	}

	if len(filteredMatches) > 1 {
		return 0, fmt.Errorf("found more than one match: %+v", filteredMatches)
	}

	return 0, fmt.Errorf("couldn't find match")
}

func AOC2022152(input string) (string, error) {
	res, err := AOC2022152Helper(input, 0, 4000000)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", res), nil
}
