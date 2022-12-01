package parseutil

import (
	"bufio"
	"bytes"
	"os"
	"strconv"

	"github.com/barszcz/advent-2022/internal/util"
)

func ReadBuffer(filename string) *bytes.Buffer {
	f := util.Must(os.ReadFile(filename))
	return bytes.NewBuffer(f)
}

func ReadInts(filename string) []int {
	f := util.Must(os.Open(filename))
	defer f.Close()
	sc := bufio.NewScanner(f)

	var res []int

	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			continue
		}
		res = append(res, n)
	}

	return res
}
