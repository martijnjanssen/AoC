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
	yPos, xPos := -1, -1
	direction := '.'
	grid := [][]rune{}
	steps := [][]rune{}
	helper.ReadLines(buf, func(l string) {
		rs := make([]rune, len(l))
		ss := make([]rune, len(l))
		for x := 0; x < len(l); x++ {
			if l[x] == '^' {
				xPos = x
				yPos = len(grid)
				direction = '^'
				rs[x] = '.'
				ss[x] = '.'
				continue
			}

			rs[x] = rune(l[x])
			ss[x] = rune(l[x])
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
			if grid[y][x] == '.' && steps[y][x] == 'X' && !(y == yPos-1 && x == xPos) {
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
	moves := map[int]bool{}
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

		// printGrid(grid, steps, yPos, xPos, direction)

		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			if grid[ny][nx] == '#' || grid[ny][nx] == 'O' {
				direction = nd
				continue
			}
			move := yPos*10000 + xPos*10 + int(direction)
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

func printGrid(grid [][]rune, steps [][]rune, yPos int, xPos int, direction rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if y == yPos && x == xPos {
				fmt.Printf("%c", direction)
			} else if steps[y][x] == 'X' {
				fmt.Printf("%c", steps[y][x])
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
