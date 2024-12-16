package day_16

import (
	"bufio"
	"container/heap"
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

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
	pq := PriorityQueue{&bot{y: sy, x: sx, direction: '>', score: 0, scores: botScores}, &bot{y: sy, x: sx, direction: '^', score: 1000, scores: botScores}}
	heap.Init(&pq)

	for pq.Len() > 0 {
		// printGrid(grid, pq, sy, sx, ey, ex)
		// _, key, _ := keyboard.GetKey()
		// switch key {
		// case keyboard.KeyEsc:
		// 	return
		// case keyboard.KeyEnter:
		// }

		rd := heap.Pop(&pq).(*bot)

		// fmt.Printf("%s %d,%d %d\n", string(rd.direction), rd.y, rd.x, rd.score)
		dy, dx := getDyDx(rd.direction)
		nScore, ny, nx := rd.score+1, rd.y+dy, rd.x+dx
		score, ok := rd.scores[getKey(ny, nx)]
		if ok && score < nScore {
			continue // Already visited this node, score is better
		}
		if ok && score == nScore {
			continue // Already visited this node, path is equal
		}
		switch grid[ny][nx] {
		case '#':
			continue
		case 'E':
			rd.scores[getKey(ny, nx)] = nScore
			pq.Update(rd, nScore)
			continue
		case '.':
			rd.scores[getKey(ny, nx)] = nScore
			pq.Update(rd, nScore)
			heap.Push(&pq, &bot{y: ny, x: nx, direction: rd.direction, score: nScore, scores: rd.scores})
		}

		switch rd.direction {
		case '^':
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '<', score: nScore + 1000, scores: rd.scores})
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '>', score: nScore + 1000, scores: rd.scores})
		case '>':
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '^', score: nScore + 1000, scores: rd.scores})
			heap.Push(&pq, &bot{y: ny, x: nx, direction: 'v', score: nScore + 1000, scores: rd.scores})
		case '<':
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '^', score: nScore + 1000, scores: rd.scores})
			heap.Push(&pq, &bot{y: ny, x: nx, direction: 'v', score: nScore + 1000, scores: rd.scores})
		case 'v':
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '<', score: nScore + 1000, scores: rd.scores})
			heap.Push(&pq, &bot{y: ny, x: nx, direction: '>', score: nScore + 1000, scores: rd.scores})
		}
	}
	a = botScores[getKey(ey, ex)]

	printGrid(grid, pq, sy, sx, ey, ex)

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

func printGrid(grid [][]rune, bots PriorityQueue, sy, sx, ey, ex int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			hasBot := 0
			for i := 0; i < len(bots); i++ {
				if bots[i].y == y && bots[i].x == x {
					hasBot++
				}
			}
			if hasBot > 0 {
				fmt.Printf("%d", hasBot)
			} else if sy == y && sx == x {
				fmt.Printf("S")
			} else if ey == y && ex == x {
				fmt.Printf("E")
			} else {
				fmt.Printf("%s", string(grid[y][x]))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
