package day_5

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type point struct {
	x int
	y int
}

type line struct {
	a *point
	b *point
}

func makeRange(a, b int) []int {
	c := 1
	if a > b {
		c = -1
	}
	r := make([]int, abs(a-b)+1)
	for i := range r {
		r[i] = a + c*i
	}
	return r
}

func (l *line) getPoints() []int {
	ps := []int{}
	if l.a.y == l.b.y {
		for _, i := range makeRange(l.a.x, l.b.x) {
			ps = append(ps, i, l.a.y)
		}
	} else if l.a.x == l.b.x {
		for _, i := range makeRange(l.a.y, l.b.y) {
			ps = append(ps, l.a.x, i)
		}
	}
	return ps
}

func (l *line) getDiagonalPoints() []int {
	ps := []int{}
	if l.a.y != l.b.y && l.a.x != l.b.x {
		ys := makeRange(l.a.y, l.b.y)
		xs := makeRange(l.a.x, l.b.x)
		for i := range ys {
			ps = append(ps, xs[i], ys[i])
		}
	}
	return ps
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
		lines = append(lines, &line{&point{ps[0], ps[1]}, &point{ps[2], ps[3]}})
		if m := max(max(ps[0], ps[1]), max(ps[2], ps[3])); m > maxV {
			maxV = m
		}
	})

	// Make grid
	grid := make([][]int, maxV+1)
	for i := range grid {
		grid[i] = make([]int, maxV+1)
	}

	// Write add non-diagonal lines to grid
	for _, l := range lines {
		points := l.getPoints()
		for i := 0; i < len(points)/2; i++ {
			grid[points[i*2]][points[i*2+1]] += 1
		}
	}
	a = calculateCriticals(grid)

	for _, l := range lines {
		points := l.getDiagonalPoints()
		for i := 0; i < len(points)/2; i++ {
			grid[points[i*2]][points[i*2+1]] += 1
		}
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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
