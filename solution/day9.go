package solution

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[8] = new(Day9)
}

type Day9 struct{}

var day9directions = map[string]complex128{
	"U": 1i,
	"D": -1i,
	"R": 1,
	"L": -1,
}

func (d *Day9) updateTailPosition(headPosition, tailPosition complex128) complex128 {
	distance := headPosition - tailPosition
	xDist := real(distance)
	yDist := imag(distance)
	newPosition := headPosition
	if xDist == 2 {
		newPosition -= 1
	}
	if xDist == -2 {
		newPosition += 1
	}
	if yDist == 2 {
		newPosition -= 1i
	}
	if yDist == -2 {
		newPosition += 1i
	}
	if newPosition != headPosition {
		return newPosition
	}
	return tailPosition
}

func (d *Day9) solve(input []byte, numKnots int) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var knotPositions []complex128
	for i := 0; i < numKnots; i++ {
		knotPositions = append(knotPositions, 0)
	}
	tailVisited := util.NewSet[complex128]()
	tailVisited.Add(0)
	for lines.Scan() {
		instruction := strings.Split(lines.Text(), " ")
		direction := instruction[0]
		steps := Must(strconv.Atoi(instruction[1]))
		for i := 0; i < steps; i++ {
			knotPositions[0] += day9directions[direction]
			for j := 1; j < numKnots; j++ {
				knotPositions[j] = d.updateTailPosition(knotPositions[j-1], knotPositions[j])
			}
			tailVisited.Add(knotPositions[numKnots-1])
		}
	}
	return strconv.Itoa(tailVisited.Size())
}

func (d *Day9) Part1(input []byte) string {
	return d.solve(input, 2)
}

func (d *Day9) Part2(input []byte) string {
	return d.solve(input, 10)
}
