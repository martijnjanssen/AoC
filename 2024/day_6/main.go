package day_6

import (
	"bufio"
	"fmt"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	xPos, yPos := -1, -1
	direction := '.'
	grid := [][]rune{}
	steps := [][]rune{}
	helper.ReadLines(buf, func(l string) {
		rs := []rune{}
		ss := []rune{}
		for x := 0; x < len(l); x++ {
			if l[x] == '^' {
				xPos = x
				yPos = len(grid)
				direction = '^'
				rs = append(rs, '.')
				ss = append(ss, '.')
				continue
			}
			rs = append(rs, rune(l[x]))
			ss = append(ss, rune(l[x]))
		}
		grid = append(grid, rs)
		steps = append(steps, ss)
	})

	hasLoop := walk(grid, steps, yPos, xPos, direction)
	if hasLoop {
		panic("has loop")
	}
	a = count(steps)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '.' && !(y == yPos-1 && x == xPos) {
				grid[y][x] = 'O'
				hasLoop := walk(grid, steps, yPos, xPos, direction)
				if hasLoop {
					b++
				}
				grid[y][x] = '.'
			}
		}
	}

	return
}

func walk(grid [][]rune, steps [][]rune, yPos int, xPos int, direction rune) bool {
	moves := map[string]bool{}
	for yPos >= 0 && yPos < len(grid) && xPos >= 0 && xPos < len(grid[yPos]) {
		ny, nx := yPos, xPos
		nd := direction
		switch direction {
		case '^':
			ny += -1
			nd = '>'
		case '<':
			nx += -1
			nd = '^'
		case '>':
			nx += 1
			nd = 'v'
		case 'v':
			ny += 1
			nd = '<'
		}

		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			if grid[ny][nx] == '#' || grid[ny][nx] == 'O' {
				direction = nd
				continue
			}
			move := fmt.Sprintf("%d,%d|%d,%d", yPos, xPos, ny, nx)
			if !moves[move] {
				moves[move] = true
			} else {
				return true
			}
			steps[ny][nx] = 'X'
			yPos, xPos = ny, nx
		} else {
			steps[yPos][xPos] = 'X'
			break
		}
	}

	return false
}

func count(grid [][]rune) int {
	res := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '.' && grid[y][x] != '#' {
				res++
			}
		}
	}
	return res
}

func printGrid(grid [][]rune, xPos int, yPos int, direction rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if y == yPos && x == xPos {
				fmt.Printf("%s", string(direction))
			} else {
				fmt.Printf("%s", string(grid[y][x]))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
