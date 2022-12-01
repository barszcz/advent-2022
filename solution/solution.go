package solution

import (
	"embed"
)

//go:embed input/*
var Inputs embed.FS

type Solution interface {
	Part1(input []byte) string
	Part2(input []byte) string
}

var Solutions = make([]Solution, 25)
