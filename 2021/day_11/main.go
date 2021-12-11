package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

type pos struct {
	r int
	c int
	e int
}

var grid = [][]pos{}

func main() {
	defer helper.Time()()

	helper.DownloadAndRead(11, func(l string) {
		row := []pos{}
		for _, s := range strings.Split(l, "") {
			i, _ := strconv.Atoi(s)
			row = append(row, pos{len(grid), len(row), i})
		}
		grid = append(grid, row)
	})

	flashes := 0
	var prevFlashes int
	for i := 0; true; i++ {
		prevFlashes = flashes
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				s := &grid[r][c]
				s.e += 1
			}
		}

		toFlash := []*pos{}
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				if grid[r][c].e > 9 {
					toFlash = append(toFlash, &grid[r][c])
				}
			}
		}

		for len(toFlash) > 0 {
			f := toFlash[0]
			toFlash = toFlash[1:]
			if f == nil || f.e == 0 || f.e <= 9 {
				continue
			}

			grid[f.r][f.c].e = 0
			flashes++
			ns := getNeighbors(f.r, f.c)
			for ni := range ns {
				if ns[ni] != nil && ns[ni].e != 0 {
					ns[ni].e += 1
					toFlash = append(toFlash, ns[ni])
				}
			}
		}

		if i == 99 {
			fmt.Printf("Answer is: %d\n", flashes)
		}
		if flashes-prevFlashes == len(grid)*len(grid[0]) {
			fmt.Printf("Answer is: %d\n", i+1)
			break
		}
	}
}

func getNeighbors(r, c int) []*pos {
	return []*pos{
		getPos(r-1, c-1), getPos(r-1, c), getPos(r-1, c+1),
		getPos(r, c-1), getPos(r, c+1),
		getPos(r+1, c-1), getPos(r+1, c), getPos(r+1, c+1),
	}
}

func getPos(r, c int) (v *pos) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
		}
	}()

	v = &grid[r][c]
	return
}
