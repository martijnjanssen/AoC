package day_17

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (a int, b int) {
	xBounds := []int{}
	yBounds := []int{}
	helper.DownloadAndRead(17, func(l string) {
		l = strings.ReplaceAll(l, "target area: x=", "")
		s := strings.Split(l, ", y=")
		xs := strings.Split(s[0], "..")
		for i, _ := range xs {
			n, _ := strconv.Atoi(xs[i])
			xBounds = append(xBounds, n)
		}
		ys := strings.Split(s[1], "..")
		for i, _ := range ys {
			n, _ := strconv.Atoi(ys[i])
			yBounds = append(yBounds, n)
		}
	})

	for svy := -85; svy < 85; svy++ {
		for svx := 22; svx < 286; svx++ {
			vx, vy := svx, svy
			x, y, max := 0, 0, 0
			overshot := false
			for !overshot {
				x += vx
				y += vy
				if y > max {
					max = y
				}
				if xBounds[0] <= x && x <= xBounds[1] && y <= yBounds[1] && yBounds[0] <= y {
					b++
					break
				}
				if xBounds[1] < x || y < yBounds[0] {
					overshot = true
				}

				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				vy--
			}

			if max > a && !overshot {
				a = max
			}
		}
	}

	return
}
