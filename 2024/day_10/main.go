package day_10

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

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	trailheads := []int{}
	grid := [][]int{}
	helper.ReadLines(buf, func(l string) {
		ps := make([]int, len(l))
		for i := 0; i < len(l); i++ {
			ss := input.SplitToInt(l, "")
			ps[i] = ss[i]

			if ps[i] == 0 {
				trailheads = append(trailheads, 1000*len(grid)+i)
			}
		}
		grid = append(grid, ps)
	})

	for _, th := range trailheads {
		q := []int{th}
		found := map[int]int{}
		score := 0
		for len(q) > 0 {
			curr := q[0]
			y, x := yx(curr)
			q = q[1:]
			point := grid[y][x]

			if point == 9 {
				found[curr]++
				score++
				continue
			}

			up := getPoint(grid, y-1, x)
			down := getPoint(grid, y+1, x)
			left := getPoint(grid, y, x-1)
			right := getPoint(grid, y, x+1)

			if up != -1 && up == point+1 {
				q = append(q, makePoint(y-1, x))
			}
			if down != -1 && down == point+1 {
				q = append(q, makePoint(y+1, x))
			}
			if left != -1 && left == point+1 {
				q = append(q, makePoint(y, x-1))
			}
			if right != -1 && right == point+1 {
				q = append(q, makePoint(y, x+1))
			}
		}
		for _, v := range found {
			a++
			b += v
		}
	}

	return
}

func getPoint(grid [][]int, y int, x int) int {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return -1
	}

	return grid[y][x]
}

func makePoint(y int, x int) int {
	return 1000*y + x
}
func yx(v int) (int, int) {
	return v / 1000, v % 1000
}
