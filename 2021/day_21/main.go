package day_21

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

type player struct {
	position int
	score    int
}

var (
	rolls int
	roll  int
	store map[[4]int][2]int
)

func (r *run) Run() (a int, b int) {
	ps := [2]player{}
	psi := 0
	rolls = 0
	roll = 1
	store = map[[4]int][2]int{}
	helper.DownloadAndRead(21, func(l string) {
		s := strings.Split(l, "starting position: ")
		n, _ := strconv.Atoi(s[1])
		ps[psi] = player{position: n - 1}
		psi++
	})

	// calculate part 2 first otherwise starting positions are modified
	b = helper.Max(play(ps[0].position, ps[1].position, 0, 0))

game:
	for {
		for i := range ps {
			p := &ps[i]
			r := 0
			for range make([]int, 3) {
				r += roll
				if roll++; roll > 100 {
					roll = 1
				}
			}
			rolls += 3
			p.position = (p.position + r) % 10
			p.score += p.position + 1
			if p.score >= 1000 {
				a = ps[1-i].score * rolls
				break game
			}
		}
	}

	return
}

func play(p1 int, p2 int, s1 int, s2 int) (int, int) {
	if s1 >= 21 {
		return 1, 0
	}
	if s2 >= 21 {
		return 0, 1
	}
	storeArr := [4]int{p1, p2, s1, s2}
	if res, ok := store[storeArr]; ok {
		return res[0], res[1]
	}
	w1, w2 := 0, 0
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				new_p1 := (p1 + d1 + d2 + d3) % 10
				new_s1 := s1 + new_p1 + 1

				wp2, wp1 := play(p2, new_p1, s2, new_s1)
				w1 += wp1
				w2 += wp2
			}
		}
	}

	store[storeArr] = [2]int{w1, w2}
	return w1, w2
}
