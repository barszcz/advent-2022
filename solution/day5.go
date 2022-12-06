package solution

import (
	"bufio"
	"bytes"
	"container/list"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	Solutions[4] = new(Day5)
}

type Day5 struct{}

var day5re = regexp.MustCompile(`move (\d+) from (\d) to (\d)`)

func (d *Day5) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	stacks := make(map[int]*list.List)
	for i := 1; i < 10; i++ {
		stacks[i] = list.New()
	}

	for lines.Scan() {
		line := lines.Text()
		// break once we've hit the bottom of the stacks
		if line[1] == '1' {
			break
		}
		for i := 0; i < 9; i++ {
			char := line[4*i+1]
			if char == ' ' {
				continue
			}
			stacks[i+1].PushBack(char)
		}
	}
	lines.Scan() // blank line

	for lines.Scan() {
		line := lines.Text()
		matches := day5re.FindStringSubmatch(line)
		count := Must(strconv.Atoi(matches[1]))
		source := Must(strconv.Atoi(matches[2]))
		destination := Must(strconv.Atoi(matches[3]))
		for i := 0; i < count; i++ {
			sourceStack := stacks[source]
			destStack := stacks[destination]
			el := sourceStack.Remove(sourceStack.Front())
			destStack.PushFront(el)
		}
	}

	var ret strings.Builder
	for i := 1; i < 10; i++ {
		char := stacks[i].Front().Value
		ret.WriteByte(char.(byte))
	}
	return ret.String()
}

func (d *Day5) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))

	stacks := make(map[int]*list.List)
	for i := 1; i < 10; i++ {
		stacks[i] = list.New()
	}

	for lines.Scan() {
		line := lines.Text()
		// break once we've hit the bottom of the stacks
		if line[1] == '1' {
			break
		}
		for i := 0; i < 9; i++ {
			char := line[4*i+1]
			if char == ' ' {
				continue
			}
			stacks[i+1].PushBack(char)
		}
	}
	lines.Scan() // blank line

	for lines.Scan() {
		line := lines.Text()
		matches := day5re.FindStringSubmatch(line)
		count := Must(strconv.Atoi(matches[1]))
		source := Must(strconv.Atoi(matches[2]))
		destination := Must(strconv.Atoi(matches[3]))
		sourceStack := stacks[source]
		destStack := stacks[destination]
		tempStack := list.New()
		for i := 0; i < count; i++ {
			el := sourceStack.Remove(sourceStack.Front())
			tempStack.PushBack(el)
		}
		destStack.PushFrontList(tempStack)
	}

	var ret strings.Builder
	for i := 1; i < 10; i++ {
		char := stacks[i].Front().Value
		ret.WriteByte(char.(byte))
	}
	return ret.String()
}
