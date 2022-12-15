package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type AOC202215MapPoint int

const (
	MPBeacon AOC202215MapPoint = iota
	MPSensor
	MPNoBeacon
)

type AOC202215SensorBeacon struct {
	Sensor   []int
	Beacon   []int
	Distance int
}

type AOC202215Map struct {
	SensorBeacons []*AOC202215SensorBeacon
	Map           map[int]map[int]AOC202215MapPoint
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
