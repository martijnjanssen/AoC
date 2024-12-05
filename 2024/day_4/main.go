package day_4

import (
	"bufio"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var (
	xmas = "XMAS"
	samx = "SAMX"
	mas  = "MAS"
	sam  = "SAM"
)

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	grid := [][]byte{}
	helper.ReadLines(buf, func(l string) {
		grid = append(grid, []byte(l))
	})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			// Solve B
			if i+2 < len(grid) && j+2 < len(grid[i]) {
				dl := string([]byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2]})
				dr := string([]byte{grid[i][j+2], grid[i+1][j+1], grid[i+2][j]})
				if (dl == mas || dl == sam) && (dr == mas || dr == sam) {
					b++
				}
			}

			// Solve A
			if i+3 < len(grid) {
				hor := string([]byte{grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j]})
				if hor == xmas || hor == samx {
					a++
				}
			}
			if j+3 < len(grid[i]) {
				ver := string([]byte{grid[i][j], grid[i][j+1], grid[i][j+2], grid[i][j+3]})
				if ver == xmas || ver == samx {
					a++
				}
			}
			if i+3 < len(grid) && j+3 < len(grid[i]) {
				diagl := string([]byte{grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]})
				if diagl == xmas || diagl == samx {
					a++
				}
				diagr := string([]byte{grid[i][j+3], grid[i+1][j+2], grid[i+2][j+1], grid[i+3][j]})
				if diagr == xmas || diagr == samx {
					a++
				}
			}
		}
	}

	return
}
