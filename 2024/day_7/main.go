package day_7

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	helper.ReadLines(buf, func(l string) {
		res, rest, _ := strings.Cut(l, ": ")
		goal, _ := strconv.Atoi(res)
		ns := input.SplitToInt(rest, " ")
		if choose(0, ns, goal, false) {
			a += goal
		} else if choose(0, ns, goal, true) {
			b += goal
		}
	})

	b = a + b

	return
}

func choose(curr int, ns []int, goal int, allowConcat bool) bool {
	if curr > goal {
		return false
	}
	if len(ns) == 0 {
		return curr == goal
	}

	mul := choose(curr*ns[0], ns[1:], goal, allowConcat)
	add := choose(curr+ns[0], ns[1:], goal, allowConcat)
	if allowConcat && !mul && !add {
		cc := curr*(int(math.Pow(10, math.Floor(math.Log10(float64(ns[0])))+1))) + ns[0]
		concat := choose(cc, ns[1:], goal, allowConcat)
		return concat
	}
	return mul || add
}
