package solution

import (
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[5] = new(Day6)
}

type Day6 struct{}

func (d *Day6) solveWithSets(input []byte, distinctCount int) string {
	for i := range input {
		if i < distinctCount {
			continue
		}
		charSet := util.NewSetFromSlice(input[i-distinctCount : i])
		if charSet.Size() == distinctCount {
			return strconv.Itoa(i)
		}
	}
	panic("unreachable")
}

func (d *Day6) solveWithCountMap(input []byte, distinctCount int) string {
	countMap := make(map[byte]int)
	for i, char := range input {
		_, ok := countMap[char]
		if !ok {
			countMap[char] = 0
		}
		countMap[char]++
		if i < distinctCount {
			continue
		}
		outChar := input[i-distinctCount]
		countMap[outChar]--
		if countMap[outChar] == 0 {
			delete(countMap, outChar)
		}
		if len(countMap) == distinctCount {
			return strconv.Itoa(i + 1)
		}
	}
	panic("unreachable")
}

func (d *Day6) Part1(input []byte) string {
	return d.solveWithCountMap(input, 4)
}

func (d *Day6) Part2(input []byte) string {
	return d.solveWithCountMap(input, 14)
}
