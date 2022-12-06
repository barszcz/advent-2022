package solution

import (
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[5] = new(Day6)
}

type Day6 struct{}

func (d *Day6) solve(input []byte, distinctCount int) string {
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

func (d *Day6) Part1(input []byte) string {
	return d.solve(input, 4)
}

func (d *Day6) Part2(input []byte) string {
	return d.solve(input, 14)
}
