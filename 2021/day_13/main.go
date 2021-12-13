package day_13

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

func count(dots [][]int) int {
	ds := map[string]bool{}
	for i := range dots {
		ds[strconv.Itoa(dots[i][0])+","+strconv.Itoa(dots[i][1])] = true
	}
	return len(ds)
}

func mirror(direction string, m int, d []int) {
	if direction == "y" {
		if d[0] > m {
			d[0] = 2*m - d[0]
		}
	} else {
		if d[1] > m {
			d[1] = 2*m - d[1]
		}
	}
}

func (r *run) Run() (a int, b int) {
	dots := [][]int{}
	mY := 0
	mX := 0
	folds := [][]string{}
	helper.DownloadAndRead(13, func(l string) {
		if l == "" {
			return
		} else if strings.Contains(l, "fold along ") {
			l = strings.ReplaceAll(l, "fold along ", "")
			f := strings.Split(l, "=")
			folds = append(folds, f)
			m, _ := strconv.Atoi(f[1])
			if f[0] == "y" {
				mY = m
			} else {
				mX = m
			}

			return
		}

		spl := strings.Split(l, ",")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])
		dots = append(dots, []int{y, x})
	})

	grid := make([][]int, mY)
	for i := range grid {
		grid[i] = make([]int, mX)
	}

	for i := range folds {
		f := folds[i]
		for j := range dots {
			d := dots[j]
			m, _ := strconv.Atoi(f[1])
			mirror(f[0], m, d)
		}
		if i == 0 {
			a = count(dots)
		}
	}
	b = count(dots)

	// for i := range dots {
	// 	d := dots[i]
	// 	grid[d[0]][d[1]] = 1
	// }

	// for i := range make([]int, mY) {
	// 	for j := range make([]int, mX) {
	// 		s := ""
	// 		if grid[i][j] == 1 {
	// 			s += "#"
	// 		} else {
	// 			s += "."
	// 		}
	// 		fmt.Printf(s)
	// 	}
	// 	fmt.Printf("\n")
	// }

	return
}
