package day_8

import (
	"bufio"
	"strconv"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	grid := [][]rune{}
	antennas := map[rune][]string{}
	antinodesA := [][]bool{}
	antinodesB := [][]bool{}
	helper.ReadLines(buf, func(l string) {
		rs := []rune{}
		for i := 0; i < len(l); i++ {
			if l[i] != '.' {
				p, ok := antennas[rune(l[i])]
				if !ok {
					p = []string{}
				}
				antennas[rune(l[i])] = append(p, strconv.Itoa(len(grid))+","+strconv.Itoa(i))
			}
			rs = append(rs, rune(l[i]))
		}
		grid = append(grid, rs)
		antinodesA = append(antinodesA, make([]bool, len(l)))
		antinodesB = append(antinodesB, make([]bool, len(l)))
	})

	for _, as := range antennas {
		for i := 0; i < len(as); i++ {
			for j := 0; j < len(as); j++ {
				if i == j {
					continue
				}

				a := input.SplitToInt(as[i], ",")
				b := input.SplitToInt(as[j], ",")
				dy := a[0] - b[0]
				dx := a[1] - b[1]

				addAntinode(antinodesA, a[0]+dy, a[1]+dx)
				addAntinode(antinodesA, b[0]+-1*dy, b[1]+-1*dx)

				addAntinodes(antinodesB, a[0], a[1], dy, dx)
				addAntinodes(antinodesB, b[0], b[1], dy*-1, dx*-1)
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if antinodesA[y][x] {
				a++
			}
			if antinodesB[y][x] {
				b++
			}
		}
	}

	return
}

func addAntinode(antinodes [][]bool, y int, x int) bool {
	if y < 0 || y >= len(antinodes) || x < 0 || x >= len(antinodes[y]) {
		return false
	}
	antinodes[y][x] = true
	return true
}

func addAntinodes(antinodes [][]bool, y int, x int, dy int, dx int) {
	for addAntinode(antinodes, y, x) {
		y += dy
		x += dx
	}
}
