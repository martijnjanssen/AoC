package day_15

import (
	"bufio"
	"fmt"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	gridScanning := true
	grid := [][]rune{}
	moves := ""
	yPos, xPos := -1, -1
	helper.ReadLines(buf, func(l string) {
		if l == "" {
			gridScanning = false
			return
		}
		if gridScanning {
			grid = append(grid, input.SplitToRune(l))
			for x := 0; x < len(grid[len(grid)-1]); x++ {
				if grid[len(grid)-1][x] == '@' {
					yPos = len(grid) - 1
					xPos = x
					grid[len(grid)-1][x] = '.'
				}
			}
			return
		}

		moves += l
	})

	for i := range moves {
		fmt.Printf("Move %s:\n", string(moves[i]))
		dy, dx := getDyDx(rune(moves[i]))
		ny, nx := yPos+dy, xPos+dx
		next := grid[ny][nx]
		nextFree := next != '#'
		if next == 'O' {
			nextFree = moveBoxes(grid, yPos+dy, xPos+dx, rune(moves[i]))
		}
		if nextFree {
			yPos, xPos = ny, nx
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				a += 100*y + x
			}
		}
	}

	printGrid(grid, yPos, xPos)

	return
}

func getDyDx(move rune) (int, int) {
	switch move {
	case '^':
		return -1, 0
	case '>':
		return 0, 1
	case '<':
		return 0, -1
	case 'v':
		return 1, 0
	default:
		return 0, 0
	}
}

func moveBoxes(grid [][]rune, yPos int, xPos int, move rune) bool {
	if grid[yPos][xPos] == '#' {
		return false
	}
	if grid[yPos][xPos] == '.' {
		return true
	}

	dy, dx := getDyDx(move)
	if moveBoxes(grid, yPos+dy, xPos+dx, move) {
		grid[yPos+dy][xPos+dx] = 'O'
		grid[yPos][xPos] = '.'
		return true
	}
	return false
}

func printGrid(grid [][]rune, yPos int, xPos int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if yPos == y && xPos == x {
				fmt.Printf("@")
			} else {
				fmt.Printf("%s", string(grid[y][x]))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
