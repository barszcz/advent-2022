package main

import (
	"fmt"

	"github.com/barszcz/advent-2022/solution"
)

func main() {

	for i, soln := range solution.Solutions {
		if soln == nil {
			continue
		}
		day := i + 1
		filename := fmt.Sprintf("input/day%d.txt", day)
		input := solution.Must(solution.Inputs.ReadFile(filename))
		fmt.Printf("day %d part 1 solution is: %s\n", day, soln.Part1(input))
		fmt.Printf("day %d part 2 solution is: %s\n", day, soln.Part2(input))
	}
}
