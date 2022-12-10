package solution

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
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
		fmt.Printf("signal: %d * %d = %d\n", i, xValues[i-1], signal)
		signalSum += signal
	}
	for i := 215; i < 225; i++ {
		fmt.Printf("%d is %d\n", i, xValues[i-1])
	}
	// fmt.Println(xValues[215:225])
	return strconv.Itoa(signalSum)
}

func (d *Day10) Part2(input []byte) string {
	return ""
}
