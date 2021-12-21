package day_21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
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
	rolls = 0
	roll  = 1
)

func (r *run) Run() (a int, b int) {
	ps := [2]player{}
	psi := 0
	helper.DownloadAndRead(21, func(l string) {
		s := strings.Split(l, "starting position: ")
		n, _ := strconv.Atoi(s[1])
		ps[psi] = player{position: n}
		psi++
	})

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
			p.position = (p.position+r-1)%10 + 1
			p.score += p.position
			fmt.Printf("player %d (score: %d): rolled %d, is in %d\n", i, p.score, r, p.position)

			if p.score >= 1000 {
				fmt.Println(ps[1-i].score)
				fmt.Println(rolls)
				a = ps[1-i].score * rolls
				break game
			}
		}
	}

	return
}
