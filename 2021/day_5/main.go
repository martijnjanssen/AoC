package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
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

func (l *line) getPoints() []*point {
	ps := []*point{}
	if l.a.y == l.b.y {
		for _, i := range makeRange(l.a.x, l.b.x) {
			ps = append(ps, &point{i, l.a.y})
		}
	} else if l.a.x == l.b.x {
		for _, i := range makeRange(l.a.y, l.b.y) {
			ps = append(ps, &point{l.a.x, i})
		}
	}
	return ps
}

func (l *line) getDiagonalPoints() []*point {
	ps := []*point{}
	if l.a.y != l.b.y && l.a.x != l.b.x {
		ys := makeRange(l.a.y, l.b.y)
		xs := makeRange(l.a.x, l.b.x)
		for i := range ys {
			ps = append(ps, &point{xs[i], ys[i]})
		}
	}
	return ps
}

func main() {
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
		for _, p := range l.getPoints() {
			grid[p.y][p.x] += 1
		}
	}
	fmt.Printf("Answer is: %d\n", calculateCriticals(grid))

	for _, l := range lines {
		for _, p := range l.getDiagonalPoints() {
			grid[p.y][p.x] += 1
		}
	}
	fmt.Printf("Answer is: %d\n", calculateCriticals(grid))
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
