package day_5

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type line struct {
	ay int
	ax int
	by int
	bx int
}

func plotStraightLine(grid [][]int, l *line) {
	if l.ay == l.by {
		if l.ax > l.bx {
			l.ax, l.bx = l.bx, l.ax
		}
		for i := l.ax; i <= l.bx; i++ {
			grid[l.ay][i]++
		}
	} else if l.ax == l.bx {
		if l.ay > l.by {
			l.ay, l.by = l.by, l.ay
		}
		for i := l.ay; i <= l.by; i++ {
			grid[i][l.ax]++
		}
	}
}

func plotDiagonalLine(grid [][]int, l *line) {
	if l.ay == l.by || l.ax == l.bx {
		return
	}

	if l.ay > l.by {
		l.ay, l.ax, l.by, l.bx = l.by, l.bx, l.ay, l.ax
	}
	d := 1
	if l.ax > l.bx {
		d = -1
	}

	for i := l.ay; i <= l.by; i++ {
		grid[i][l.ax]++
		l.ax += d
	}
}

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (a int, b int) {
	lines := []*line{}
	maxV := 0

	helper.DownloadAndRead(5, func(l string) {
		ps := parseLine(l)
		lines = append(lines, &line{ps[0], ps[1], ps[2], ps[3]})
		if m := helper.Max(helper.Max(ps[0], ps[1]), helper.Max(ps[2], ps[3])); m > maxV {
			maxV = m
		}
	})

	// Make grid
	grid := make([][]int, maxV+1)
	for i := range grid {
		grid[i] = make([]int, maxV+1)
	}

	// Write add non-diagonal lines to grid
	for i := range lines {
		plotStraightLine(grid, lines[i])
	}
	a = calculateCriticals(grid)

	for i := range lines {
		plotDiagonalLine(grid, lines[i])
	}
	b = calculateCriticals(grid)
	return
}

func calculateCriticals(grid [][]int) int {
	c := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] > 1 {
				c++
			}
		}
	}
	return c
}

func parseLine(l string) []int {
	l = strings.ReplaceAll(l, " -> ", ",")
	ss := strings.Split(l, ",")

	r := []int{}
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		r = append(r, i)
	}
	return r
}
