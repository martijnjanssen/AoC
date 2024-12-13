package day_13

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var aX, aY = -1, -1
var bX, bY = -1, -1
var scoreCache = map[string]int{}

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

		_, afterX, _ := strings.Cut(l, "X=")
		xString, afterY, _ := strings.Cut(afterX, ",")
		goalX, _ := strconv.Atoi(xString)
		_, yString, _ := strings.Cut(afterY, "Y=")
		goalY, _ := strconv.Atoi(yString)

		choice(goalX, goalY, 0, 0, 0, 0, true)
		a += scoreCache[cacheKey(goalX, goalY)]

		if scoreCache[cacheKey(goalX, goalY)] == 0 {
			scoreCache = make(map[string]int, len(scoreCache))
			goalX += 10000000000000
			goalY += 10000000000000
			choice(goalX, goalY, 0, 0, 0, 0, false)
			b += scoreCache[cacheKey(goalX, goalY)]
			fmt.Printf("%s\n(%d, %d):   a(%d, %d)   b(%d, %d)\n\n", l, goalX, goalY, aX, aY, bX, bY)
		}

		aX, aY, bX, bY = -1, -1, -1, -1
		scoreCache = make(map[string]int, len(scoreCache))
	})

	return
}

func cacheKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
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
	_, afterX, _ := strings.Cut(l, "X+")
	xString, afterY, _ := strings.Cut(afterX, ",")
	x, _ := strconv.Atoi(xString)
	_, yString, _ := strings.Cut(afterY, "Y+")
	y, _ := strconv.Atoi(yString)

	return x, y
}
