package day_16

import (
	"bufio"
	"container/heap"
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
	grid := [][]rune{}
	sy, sx, ey, ex := -1, -1, -1, -1
	helper.ReadLines(buf, func(l string) {
		ls := input.SplitToRune(l)

		for i := 0; i < len(ls); i++ {
			if ls[i] == 'E' {
				ey = len(grid)
				ex = i
			}
			if ls[i] == 'S' {
				sy = len(grid)
				sx = i
			}
		}

		grid = append(grid, ls)
	})

	botScores := map[int]int{}
	pq := PriorityQueue{&bot{y: sy, x: sx, direction: '>', score: 0, visited: fmt.Sprintf("%d,%d", sy, sx)}}
	finished := []*bot{}
	heap.Init(&pq)

	for pq.Len() > 0 {
		rd := heap.Pop(&pq).(*bot)

		if grid[rd.y][rd.x] == 'E' {
			finished = append(finished, rd)
			continue
		}

		for _, d := range []rune{'^', 'v', '<', '>'} {
			dy, dx := getDyDx(d)
			if grid[rd.y+dy][rd.x+dx] == '#' {
				continue
			}

			cdy, cdx := getDyDx(rd.direction)
			if cdy*-1 == dy && cdx*-1 == dx {
				continue
			}
			nScore := rd.score + 1
			if d != rd.direction {
				nScore += 1000
				botScores[getKey(rd.y, rd.x)] = nScore
			}
			score, ok := botScores[getKey(rd.y+dy, rd.x+dx)]
			if ok && score < rd.score {
				continue // Already visited this node, score is better
			}
			botScores[getKey(rd.y+dy, rd.x+dx)] = nScore
			v := fmt.Sprintf("%s|%d,%d", rd.visited, rd.y+dy, rd.x+dx)

			heap.Push(&pq, &bot{y: rd.y + dy, x: rd.x + dx, direction: d, score: nScore, visited: v})

		}
	}
	a = botScores[getKey(ey, ex)]

	count := map[string]bool{}
	for _, rd := range finished {
		if rd.score > botScores[getKey(ey, ex)] {
			continue
		}
		for _, v := range strings.Split(rd.visited, "|") {
			count[v] = true
		}
	}

	b = len(count)

	return
}

func getKey(y int, x int) int {
	return y*10000 + x
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
