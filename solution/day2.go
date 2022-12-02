package solution

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

func init() {
	Solutions[1] = new(Day2)
}

type Day2 struct{}

type rpsMove int

const (
	rock = iota + 1
	paper
	scissors
)

type rpsResult int

const (
	loss = iota * 3
	draw
	win
)

func (d *Day2) playerResult(playerMove, opponentMove rpsMove) rpsResult {
	if playerMove == opponentMove {
		return draw
	}
	if playerMove == rock && opponentMove == paper {
		return loss
	}
	if playerMove == rock && opponentMove == scissors {
		return win
	}
	if playerMove == paper && opponentMove == rock {
		return win
	}
	if playerMove == paper && opponentMove == scissors {
		return loss
	}
	if playerMove == scissors && opponentMove == rock {
		return loss
	}
	if playerMove == scissors && opponentMove == paper {
		return win
	}
	panic(errors.New("unreachable"))
}

var movesParseDict = map[string]rpsMove{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

func (d *Day2) Part1(input []byte) string {
	lines := bufio.NewScanner(bytes.NewReader(input))
	var score int
	for lines.Scan() {
		line := lines.Text()
		rawMoves := strings.Split(line, " ")
		opponentMove := movesParseDict[rawMoves[0]]
		playerMove := movesParseDict[rawMoves[1]]
		score += int(playerMove) + int(d.playerResult(playerMove, opponentMove))
	}
	return strconv.Itoa(score)
}

func (d *Day2) Part2(input []byte) string {
	return ""
}
