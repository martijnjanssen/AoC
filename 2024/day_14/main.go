package day_14

import (
	"bufio"
	"fmt"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

type bot struct {
	py int
	px int
	dy int
	dx int

	endy int
	endx int
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	xSize := 101
	ySize := 103
	bots := []*bot{}
	q1, q2, q3, q4 := 0, 0, 0, 0

	helper.ReadLines(buf, func(l string) {
		b := &bot{}
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &b.px, &b.py, &b.dx, &b.dy)
		bots = append(bots, b)
	})

	for seconds := 1; seconds < 10000; seconds++ {
		for bI := range bots {
			bots[bI].endy = (bots[bI].py + seconds*bots[bI].dy) % ySize
			bots[bI].endx = (bots[bI].px + seconds*bots[bI].dx) % xSize
			for bots[bI].endy < 0 { // Modulo does not work in go with negative numbers, so need to do this...
				bots[bI].endy += ySize
			}
			for bots[bI].endx < 0 {
				bots[bI].endx += xSize
			}

			if seconds == 100 {
				if bots[bI].endy < ySize/2 {
					if bots[bI].endx < xSize/2 {
						q1++
					} else if bots[bI].endx > xSize/2 {
						q2++
					}
				} else if bots[bI].endy > ySize/2 {
					if bots[bI].endx < xSize/2 {
						q3++
					} else if bots[bI].endx > xSize/2 {
						q4++
					}
				}
			}
		}

		if seconds == 100 {
			a = q1 * q2 * q3 * q4
		}

		positions := make(map[int]bool, len(bots))
		valid := true
		for i := range bots {
			key := bots[i].endx*1000 + bots[i].endy
			_, ok := positions[key]
			if ok {
				valid = false
				break
			}
			positions[key] = true
		}
		if valid {
			b = seconds
			break
		}
	}

	return
}
