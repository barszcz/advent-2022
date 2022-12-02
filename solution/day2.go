package solution

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/barszcz/advent-2022/internal/util"
)

func init() {
	Solutions[1] = new(Day2)
}

type Day2 struct{}

type rpsMove int

const (
	rock rpsMove = iota + 1
	paper
	scissors
)

type rpsResult int

const (
	loss rpsResult = iota * 3
	draw
	win
)

var resultMap = map[rpsMove]util.Bimap[rpsMove, rpsResult]{
	rock: util.NewBimapFromMap(map[rpsMove]rpsResult{
		rock:     draw,
		paper:    win,
		scissors: loss,
	}),
	paper: util.NewBimapFromMap(map[rpsMove]rpsResult{
		rock:     loss,
		paper:    draw,
		scissors: win,
	}),
	scissors: util.NewBimapFromMap(map[rpsMove]rpsResult{
		rock:     win,
		paper:    loss,
		scissors: draw,
	}),
}

var movesParseDict = map[string]rpsMove{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var resultsParseDict = map[string]rpsResult{
	"X": loss,
	"Y": draw,
	"Z": win,
}

func (d *Day2) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var score int
	for lines.Scan() {
		line := lines.Text()
		rawMoves := strings.Split(line, " ")
		opponentMove := movesParseDict[rawMoves[0]]
		playerMove := movesParseDict[rawMoves[1]]
		innerMap := resultMap[opponentMove]
		result, _ := innerMap.Get(playerMove)
		score += int(playerMove) + int(result)
	}
	return strconv.Itoa(score)
}

func (d *Day2) Part2(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var score int
	for lines.Scan() {
		line := lines.Text()
		rawMoves := strings.Split(line, " ")
		opponentMove := movesParseDict[rawMoves[0]]
		result := resultsParseDict[rawMoves[1]]
		innerMap := resultMap[opponentMove]
		playerMove, _ := innerMap.InverseGet(result)
		score += int(playerMove) + int(result)
	}
	return strconv.Itoa(score)
}
