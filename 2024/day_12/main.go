package day_12

import (
	"bufio"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var grid = [][]rune{}
var done = map[int]bool{}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	helper.ReadLines(buf, func(l string) {
		ls := input.SplitToRune(l)

		grid = append(grid, ls)
	})

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			s, b := count(y, x)
			a += s * b
		}

	}

	return
}

func count(y int, x int) (int, int) {
	if done[y*1000+x] {
		return 0, 0
	}
	done[y*1000+x] = true

	p := getPoint(y, x)
	up, down, left, right := getPoint(y-1, x), getPoint(y+1, x), getPoint(y, x-1), getPoint(y, x+1)

	borders := 0
	size := 1
	if p == up {
		ns, nb := count(y-1, x)
		size += ns
		borders += nb
	} else {
		borders++
	}
	if p == down {
		ns, nb := count(y+1, x)
		size += ns
		borders += nb
	} else {
		borders++
	}
	if p == left {
		ns, nb := count(y, x-1)
		size += ns
		borders += nb
	} else {
		borders++
	}
	if p == right {
		ns, nb := count(y, x+1)
		size += ns
		borders += nb
	} else {
		borders++
	}

	return size, borders
}

func getPoint(y int, x int) rune {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return '.'
	}

	return grid[y][x]
}
