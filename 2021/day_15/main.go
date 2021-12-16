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
			expanded[y*mY*5+x] = &point{y: y, x: x, r: math.MaxInt, small: y < mY && x < mX}
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

		rs := updateNeighbors(tip.y, tip.x, tip.r)
		for i := range rs {
			if rs[i] == nil {
				continue
			}
			heap.Push(pq, rs[i])
		}
	}
}

func updateNeighbors(y int, x int, r int) []*point {
	return []*point{
		updatePoint(y, x+1, r),
		updatePoint(y+1, x, r),
		updatePoint(y, x-1, r),
		updatePoint(y-1, x, r),
	}
}

func updatePoint(y int, x int, r int) *point {
	if y < 0 || y >= mY*5 || x < 0 || x >= mX*5 {
		return nil
	}
	nr := r + (cave[(y%mY)*mY+x%mX]+y/mY+x/mX-1)%9 + 1
	if nr < expanded[y*mY*5+x].r {
		expanded[y*mY*5+x].r = nr
		return expanded[y*mY*5+x]
	}
	return nil
}
