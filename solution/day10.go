package solution

import (
	"bufio"
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[9] = new(Day10)
}

type Day10 struct{}

func (d *Day10) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	x := 1
	xValues := []int{1}
	for lines.Scan() {
		line := lines.Text()
		xValues = append(xValues, x)
		if line == "noop" {
			continue
		}
		incr := Must(strconv.Atoi(strings.TrimPrefix(line, "addx ")))
		x += incr
		xValues = append(xValues, x)
	}

	signalSum := 0
	for i := 20; i < len(xValues); i += 40 {
		signal := xValues[i-1] * i
		signalSum += signal
	}
	return strconv.Itoa(signalSum)
}

func (d *Day10) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	x := 1
	xValues := []int{1}
	for lines.Scan() {
		line := lines.Text()
		xValues = append(xValues, x)
		if line == "noop" {
			continue
		}
		incr := Must(strconv.Atoi(strings.TrimPrefix(line, "addx ")))
		x += incr
		xValues = append(xValues, x)
	}

	pixels := make([]bool, 0, 240)
	for i := 0; i < 240; i++ {
		pixelVal := math.Abs(float64(((i % 40) - xValues[i]))) <= 1
		pixels = append(pixels, pixelVal)
	}

	pixelLines := util.ChunkSlice(pixels, 40)

	var builder strings.Builder

	builder.WriteString("\n")
	for _, pixelLine := range pixelLines {
		for _, pixel := range pixelLine {
			if pixel {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()

}
