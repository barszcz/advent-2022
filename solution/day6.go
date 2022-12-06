package solution

import (
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[5] = new(Day6)
}

type Day6 struct{}

func (d *Day6) Part1(input []byte) string {
	for i := range input {
		if i < 4 {
			continue
		}
		charSet := util.NewSetFromSlice(input[i-4 : i])
		if charSet.Size() == 4 {
			return strconv.Itoa(i)
		}
	}
	panic("unreachable")
}

func (d *Day6) Part2(input []byte) string {
	for i := range input {
		if i < 14 {
			continue
		}
		charSet := util.NewSetFromSlice(input[i-14 : i])
		if charSet.Size() == 14 {
			return strconv.Itoa(i)
		}
	}
	panic("unreachable")
}
