package day_15

import (
	"container/heap"
	"math"
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
	expanded []*point
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
	expanded = make([]*point, mY*mX*5*5)
	for y := 0; y < mY*5; y++ {
		for x := 0; x < mX*5; x++ {
			expanded[y*mY*5+x] = &point{y: y, x: x, r: math.MaxInt32, small: y < mY && x < mX}
		}
	}
	expanded[0].r = 0

	pq := make(PriorityQueue, 1)
	pq[0] = expanded[0]
	heap.Init(&pq)

	handleTodo(&pq, true)
	a = expanded[(mY-1)*5*mX+mX-1].r

	handleTodo(&pq, false)
	b = expanded[mY*5*mX*5-1].r

	return
}

func handleTodo(pq *PriorityQueue, small bool) {
	for pq.Len() > 0 {
		tip := heap.Pop(pq).(*point)
		if small && !tip.small {
			return
		}

		updatePoint(pq, tip.y, tip.x+1, tip.r)
		updatePoint(pq, tip.y+1, tip.x, tip.r)
		updatePoint(pq, tip.y, tip.x-1, tip.r)
		updatePoint(pq, tip.y-1, tip.x, tip.r)
	}
}

func updatePoint(pq *PriorityQueue, y int, x int, r int) {
	if y < 0 || y >= mY*5 || x < 0 || x >= mX*5 {
		return
	}

	nr := r + (cave[(y%mY)*mY+x%mX]+y/mY+x/mX-1)%9 + 1
	p := expanded[y*mY*5+x]
	if nr < p.r {
		if p.inQueue {
			pq.update(p, nr)
		} else {
			p.r = nr
			heap.Push(pq, p)
		}
	}
}
