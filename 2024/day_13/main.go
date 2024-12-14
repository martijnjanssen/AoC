package day_13

import (
	"bufio"
	"fmt"
	"math"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var aX, aY = -1, -1
var bX, bY = -1, -1
var scoreCache = map[int]int{}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	helper.ReadLines(buf, func(l string) {
		if l == "" {
			return
		}
		if aX == -1 {
			aX, aY = splitLine(l)
			return
		}
		if bX == -1 {
			bX, bY = splitLine(l)
			return
		}

		var goalX, goalY int
		fmt.Sscanf(l, "Prize: X=%d, Y=%d", &goalX, &goalY)

		choice(goalX, goalY, 0, 0, 0, 0, true)
		a += scoreCache[cacheKey(goalX, goalY)]

		if scoreCache[cacheKey(goalX, goalY)] == 0 {
			goalX += 10000000000000
			goalY += 10000000000000

			aPress := math.Round(((float64(goalX) * float64(bY) / float64(bX)) - float64(goalY)) / ((float64(bY) * float64(aX) / float64(bX)) - float64(aY)))
			bPress := math.Round(float64(goalX)-float64(aPress)*float64(aX)) / float64(bX)

			b += int(aPress)*3 + int(bPress)
		}

		aX, aY, bX, bY = -1, -1, -1, -1
		scoreCache = make(map[int]int, len(scoreCache))
	})

	return
}

func cacheKey(x int, y int) int {
	return x*1000000 + y
}

func choice(goalX int, goalY int, x int, y int, a int, b int, limitPress bool) {
	if limitPress && ((a > 100) || (b > 100)) {
		return
	}

	if x > goalX || y > goalY {
		return
	}
	newScore := a*3 + b
	key := cacheKey(x, y)
	score, ok := scoreCache[key]
	if !ok || newScore < score {
		scoreCache[key] = newScore
	} else if newScore >= score {
		return
	}
	if x == goalX && y == goalY {
		return
	}

	choice(goalX, goalY, x+aX, y+aY, a+1, b, limitPress)
	choice(goalX, goalY, x+bX, y+bY, a, b+1, limitPress)
}

func splitLine(l string) (int, int) {
	var x, y int
	fmt.Sscanf(l[10:], "X+%d, Y+%d", &x, &y)
	return x, y
}
