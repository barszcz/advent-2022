package solution

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[7] = new(Day8)
}

type Day8 struct{}

func (d *Day8) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var grid [][]int
	for lines.Scan() {
		line := lines.Text()
		var row []int
		for _, digit := range strings.Split(line, "") {
			row = append(row, Must(strconv.Atoi(digit)))
		}
		grid = append(grid, row)
	}

	gridLen := len(grid)

	seen := util.NewSet[string]()

	// look from left
	for i := 0; i < gridLen; i++ {
		maxHeightSeen := -1
		for j := 0; j < gridLen; j++ {
			treeHeight := grid[i][j]
			if treeHeight > maxHeightSeen {
				maxHeightSeen = treeHeight
				seen.Add(fmt.Sprintf("%d,%d", i, j))
				if treeHeight == 9 {
					break
				}
			}
		}
	}

	// look from right
	for i := 0; i < gridLen; i++ {
		maxHeightSeen := -1
		for j := gridLen - 1; j >= 0; j-- {
			treeHeight := grid[i][j]
			if treeHeight > maxHeightSeen {
				maxHeightSeen = treeHeight
				seen.Add(fmt.Sprintf("%d,%d", i, j))
				if treeHeight == 9 {
					break
				}
			}
		}
	}

	// look from top
	for j := 0; j < gridLen; j++ {
		maxHeightSeen := -1
		for i := 0; i < gridLen; i++ {
			treeHeight := grid[i][j]
			if treeHeight > maxHeightSeen {
				maxHeightSeen = treeHeight
				seen.Add(fmt.Sprintf("%d,%d", i, j))
				if treeHeight == 9 {
					break
				}
			}
		}
	}

	// look from bottom
	for j := 0; j < gridLen; j++ {
		maxHeightSeen := -1
		for i := gridLen - 1; i >= 0; i-- {
			treeHeight := grid[i][j]
			if treeHeight > maxHeightSeen {
				maxHeightSeen = treeHeight
				seen.Add(fmt.Sprintf("%d,%d", i, j))
				if treeHeight == 9 {
					break
				}
			}
		}
	}

	return strconv.Itoa(seen.Size())
}

func (d *Day8) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var grid [][]int
	for lines.Scan() {
		line := lines.Text()
		var row []int
		for _, digit := range strings.Split(line, "") {
			row = append(row, Must(strconv.Atoi(digit)))
		}
		grid = append(grid, row)
	}

	gridLen := len(grid)

	var scores [][]int
	for i := 0; i < gridLen; i++ {
		var row []int
		for j := 0; j < gridLen; j++ {
			row = append(row, 1)
		}
		scores = append(scores, row)
	}

	// look from left
	for i := 0; i < gridLen; i++ {
	tree1:
		for j := 0; j < gridLen; j++ {
			treeHeight := grid[i][j]
			for k := j - 1; k >= 0; k-- {
				if grid[i][k] >= treeHeight {
					scores[i][j] *= j - k
					continue tree1
				}
			}
			scores[i][j] *= j
		}
	}

	// look from right
	for i := 0; i < gridLen; i++ {
	tree2:
		for j := gridLen - 1; j >= 0; j-- {
			treeHeight := grid[i][j]
			for k := j + 1; k < gridLen; k++ {
				if grid[i][k] >= treeHeight {
					scores[i][j] *= k - j
					continue tree2
				}
			}
			scores[i][j] *= gridLen - (j + 1)
		}
	}

	// look from top
	for j := 0; j < gridLen; j++ {
	tree3:
		for i := 0; i < gridLen; i++ {
			treeHeight := grid[i][j]
			for k := i - 1; k >= 0; k-- {
				if grid[k][j] >= treeHeight {
					scores[i][j] *= i - k
					continue tree3
				}
			}
			scores[i][j] *= i
		}
	}

	// look from bottom
	for j := 0; j < gridLen; j++ {
	tree4:
		for i := gridLen - 1; i >= 0; i-- {
			treeHeight := grid[i][j]
			for k := i + 1; k < gridLen; k++ {
				if grid[k][j] >= treeHeight {
					scores[i][j] *= k - i
					continue tree4
				}
			}
			scores[i][j] *= gridLen - (i + 1)
		}
	}

	maxScore := 0
	for i := range scores {
		for j := range scores[i] {
			if scores[i][j] > maxScore {
				maxScore = scores[i][j]
			}
		}
	}

	return strconv.Itoa(maxScore)
}
