package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

var rows, cols int
var grid = [][]int{}
var basins = []int{}

func main() {
	defer helper.Time()()

	lows := 0
	helper.DownloadAndRead(9, func(l string) {
		row := []int{}
		for _, s := range strings.Split(l, "") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		grid = append(grid, row)
	})
	rows = len(grid)
	cols = len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] < lowestNeighbor(r, c) {
				lows += grid[r][c] + 1
			}
		}
	}
	fmt.Printf("Answer is: %d\n", lows)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] < lowestNeighbor(r, c) {
				basins = append(basins, findBasinSize(r, c))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Printf("Answer is: %d\n", basins[0]*basins[1]*basins[2])
}

func findBasinSize(r int, c int) int {
	if !onGrid(r, c) || grid[r][c] == 9 {
		return 0
	}

	grid[r][c] = 9
	return 1 + findBasinSize(r-1, c) + findBasinSize(r+1, c) + findBasinSize(r, c-1) + findBasinSize(r, c+1)
}

func lowestNeighbor(r int, c int) int {
	return helper.Min(helper.Min(get(r-1, c), get(r+1, c)), helper.Min(get(r, c-1), get(r, c+1)))
}

func get(r int, c int) int {
	if !onGrid(r, c) {
		return 10
	}
	return grid[r][c]
}

func onGrid(r int, c int) bool {
	if r < 0 || c < 0 || r >= rows || c >= cols {
		return false
	}
	return true
}
