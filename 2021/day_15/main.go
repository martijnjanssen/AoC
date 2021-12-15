package day_15

import (
	"container/heap"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var (
	cave     []int
	expanded []int
	mY       int
	mX       int
)

func (r *run) Run() (a int, b int) {
	cave = []int{}
	helper.DownloadAndRead(15, func(l string) {
		spl := strings.Split(l, "")
		ns := []int{}
		for i := range spl {
			n, _ := strconv.Atoi(spl[i])
			ns = append(ns, n)
		}
		cave = append(cave, ns...)

		if mX == 0 {
			mX = len(ns)
			mY = len(ns)
		}
	})
	expanded = make([]int, mY*mX*5*5)

	pq := make(PriorityQueue, 1)
	pq[0] = &point{y: 0, x: 0, r: 0, small: true}
	heap.Init(&pq)

	handleTodo(&pq, true)
	a = expanded[(mY-1)*5*mX+mX-1]

	handleTodo(&pq, false)
	b = expanded[mY*5*mX*5-1]

	return
}

func handleTodo(pq *PriorityQueue, small bool) {
	for pq.Len() > 0 {
		tip := heap.Pop(pq).(*point)
		if small && !tip.small {
			return
		}
		if expanded[tip.y*mY*5+tip.x] != 0 && tip.r >= expanded[tip.y*mY*5+tip.x] {
			continue
		}
		expanded[tip.y*mY*5+tip.x] = tip.r

		rs := getRisks(tip.y, tip.x)
		// fmt.Println(rs)
		for i := range rs {
			if rs[i] == nil {
				continue
			}
			rs[i].r += tip.r
			heap.Push(pq, rs[i])
		}
	}
}

func getRisks(y int, x int) []*point {
	return []*point{
		createPoint(y, x+1),
		createPoint(y+1, x),
		createPoint(y, x-1),
		createPoint(y-1, x),
	}
}

func createPoint(y int, x int) *point {
	if y >= 0 && y < mY*5 && x >= 0 && x < mX*5 {
		r := (cave[(y%mY)*mY+x%mX]+y/mY+x/mX-1)%9 + 1
		return &point{y: y, x: x, r: r, small: y < mY && x < mX}
	}
	return nil
}
