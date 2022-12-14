package solution

import (
	"bufio"
	"bytes"
	"container/heap"
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
	"golang.org/x/exp/constraints"
)

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func genericMax[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}

func init() {
	Solutions[0] = new(Day1)
}

type Day1 struct{}

func (d *Day1) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	var currentCalorieSum, maxCalorieSum int

	for lines.Scan() {
		line := lines.Text()
		if line == "" {
			maxCalorieSum = genericMax(maxCalorieSum, currentCalorieSum)
			currentCalorieSum = 0
			continue
		}
		currentCalorieSum += Must(strconv.Atoi(line))
	}

	return strconv.Itoa(maxCalorieSum)
}

func (d *Day1) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	var currentCalorieSum int
	calorieSums := &util.MaxHeap[int]{}

	for lines.Scan() {
		line := lines.Text()
		if line == "" {
			heap.Push(calorieSums, currentCalorieSum)
			currentCalorieSum = 0
			continue
		}
		currentCalorieSum += Must(strconv.Atoi(line))
	}

	var topThreeSum int
	for i := 0; i < 3; i++ {
		topThreeSum += heap.Pop(calorieSums).(int)
	}

	return strconv.Itoa(topThreeSum)
}
