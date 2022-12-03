package solution

import (
	"bufio"
	"bytes"
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[2] = new(Day3)
}

type Day3 struct{}

func (d *Day3) getPriority(item byte) int {
	// lowercase
	if item > 96 {
		return int(item - 96)
	}
	// uppercase
	return int(26 + item - 64)
}

func (d *Day3) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	prioritySum := 0
	for lines.Scan() {
		line := lines.Text()
		rucksackLength := len(line) / 2
		seen := util.NewSet[byte]()

		for i := 0; i < rucksackLength; i++ {
			seen.Add(line[i])
		}
		for i := rucksackLength; i < len(line); i++ {
			item := line[i]
			if seen.Has(item) {
				prioritySum += d.getPriority(item)
				break
			}
		}
	}
	return strconv.Itoa(prioritySum)
}

func (d *Day3) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	chunk := make([]util.Set[byte], 0, 3)
	prioritySum := 0
	for lines.Scan() {
		if len(chunk) == 3 {
			item := chunk[0].Intersection(chunk[1]).Intersection(chunk[2]).Slice()[0]
			prioritySum += d.getPriority(item)
			chunk = chunk[:0]
		}
		chunk = append(chunk, util.NewSetFromSlice(lines.Bytes()))
	}
	// one last time at end of loop
	item := chunk[0].Intersection(chunk[1]).Intersection(chunk[2]).Slice()[0]
	prioritySum += d.getPriority(item)
	return strconv.Itoa(prioritySum)
}
