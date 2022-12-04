package solution

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func init() {
	Solutions[3] = new(Day4)
}

type Day4 struct{}

type assignment struct {
	lowerBound int
	upperBound int
}

func (d *Day4) parseAssignment(line string) assignment {
	split := strings.Split(line, "-")
	return assignment{
		lowerBound: Must(strconv.Atoi(split[0])),
		upperBound: Must(strconv.Atoi(split[1])),
	}
}

func (d *Day4) assignmentsFullyOverlap(a1, a2 assignment) bool {
	// no overlap at all
	if !d.assignmentsOverlapAtAll(a1, a2) {
		return false
	}
	if a1.lowerBound <= a2.lowerBound && a2.upperBound <= a1.upperBound {
		return true
	}
	if a2.lowerBound <= a1.lowerBound && a1.upperBound <= a2.upperBound {
		return true
	}
	return false
}

func (d *Day4) assignmentsOverlapAtAll(a1, a2 assignment) bool {
	// no overlap at all
	if a1.lowerBound > a2.upperBound || a2.lowerBound > a1.upperBound {
		return false
	}
	return true
}

func (d *Day4) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	overlappingPairs := 0
	for lines.Scan() {
		rawAssignments := strings.Split(lines.Text(), ",")
		a1 := d.parseAssignment(rawAssignments[0])
		a2 := d.parseAssignment(rawAssignments[1])

		if d.assignmentsFullyOverlap(a1, a2) {
			overlappingPairs++
		}
	}
	return strconv.Itoa(overlappingPairs)
}

func (d *Day4) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	overlappingPairs := 0
	for lines.Scan() {
		rawAssignments := strings.Split(lines.Text(), ",")
		a1 := d.parseAssignment(rawAssignments[0])
		a2 := d.parseAssignment(rawAssignments[1])

		if d.assignmentsOverlapAtAll(a1, a2) {
			overlappingPairs++
		}
	}
	return strconv.Itoa(overlappingPairs)
}
