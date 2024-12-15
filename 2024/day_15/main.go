package day_15

import (
	"bufio"
	"fmt"
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
	// Enable to walk through the grid
	// if err := keyboard.Open(); err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	_ = keyboard.Close()
	// }()

	gridScanning := true
	grid := [][]rune{}
	expandedGrid := [][]rune{}
	moves := ""
	yPos, xPos := -1, -1
	expandedYPos, expandedXPos := -1, -1

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

			l = strings.ReplaceAll(l, "#", "##")
			l = strings.ReplaceAll(l, "O", "[]")
			l = strings.ReplaceAll(l, ".", "..")
			l = strings.ReplaceAll(l, "@", "@.")
			expandedGrid = append(expandedGrid, input.SplitToRune(l))
			return
		}

		moves += l
	})
	expandedYPos, expandedXPos = yPos, xPos*2
	expandedGrid[expandedYPos][expandedXPos] = '.'

	a = solve(grid, moves, yPos, xPos)
	b = solve(expandedGrid, moves, expandedYPos, expandedXPos)

	return
}

func solve(grid [][]rune, moves string, yPos int, xPos int) int {
	score := 0
	for i := range moves {
		move := moves[i]

		// Enable this code block to walk through the grid
		// _, key, _ := keyboard.GetKey()
		// switch key {
		// case keyboard.KeyEsc:
		// 	return 0
		// case keyboard.KeyArrowUp:
		// 	move = '^'
		// case keyboard.KeyArrowDown:
		// 	move = 'v'
		// case keyboard.KeyArrowLeft:
		// 	move = '<'
		// case keyboard.KeyArrowRight:
		// 	move = '>'
		// }

		dy, dx := getDyDx(rune(move))
		ny, nx := yPos+dy, xPos+dx
		next := grid[ny][nx]
		switch next {
		case '.':
			yPos, xPos = ny, nx
		case '[', ']', 'O':
			if canMove(grid, ny, nx, rune(move)) {
				moveBoxes(grid, ny, nx, rune(move))
				yPos, xPos = ny, nx
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '[' || grid[y][x] == 'O' {
				score += 100*y + x
			}
		}
	}

	return score
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
		panic("invalid move")
	}
}

func canMove(grid [][]rune, y int, x int, move rune) bool {
	if grid[y][x] == '#' {
		return false
	}
	if grid[y][x] == '.' {
		return true
	}

	dy, dx := getDyDx(move)
	if move == '<' || move == '>' {
		return canMove(grid, y+dy, x+dx, move)
	} else {
		if grid[y][x] == 'O' {
			return canMove(grid, y+dy, x+dx, move)
		} else if grid[y][x] == '[' {
			moveLeft := canMove(grid, y+dy, x+dx, move)
			moveRight := canMove(grid, y+dy, x+dx+1, move)
			return !(!moveLeft || !moveRight)
		} else if grid[y][x] == ']' {
			moveLeft := canMove(grid, y+dy, x+dx-1, move)
			moveRight := canMove(grid, y+dy, x+dx, move)
			return !(!moveLeft || !moveRight)
		}
	}
	return false
}

func moveBoxes(grid [][]rune, y int, x int, move rune) {
	if grid[y][x] == '#' {
		panic("should not be pushing against wall")
	}
	if grid[y][x] == '.' {
		return
	}

	dy, dx := getDyDx(move)
	moveBoxes(grid, y+dy, x+dx, move)
	if move == '^' || move == 'v' {
		switch grid[y][x] {
		case 'O':
			break
		case '[':
			moveBoxes(grid, y+dy, x+dx+1, move)
			grid[y+dy][x+dx+1] = grid[y][x+1]
			grid[y][x+1] = '.'
		case ']':
			moveBoxes(grid, y+dy, x+dx-1, move)
			grid[y+dy][x+dx-1] = grid[y][x-1]
			grid[y][x-1] = '.'
		}
	}
	grid[y+dy][x+dx] = grid[y][x]
	grid[y][x] = '.'
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
