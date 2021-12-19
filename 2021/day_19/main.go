package day_19

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

type scanner struct {
	bs [][]int
}

var (
	rotations = [][]int{
		{1, 2, 3}, {-2, 1, 3}, {-1, -2, 3}, {2, -1, 3},
		{1, -2, -3}, {-2, -1, -3}, {-1, 2, -3}, {2, 1, -3},
		{2, 3, 1}, {3, -2, 1}, {-2, -3, 1}, {-3, 2, 1},
		{2, -3, -1}, {-3, -2, -1}, {-2, 3, -1}, {3, 2, -1},
		{3, 1, 2}, {1, -3, 2}, {-3, -1, 2}, {-1, 3, 2},
		{3, -1, -2}, {-1, -3, -2}, {-3, 1, -2}, {1, 3, -2},
	}
)

func (r *run) Run() (a int, b int) {
	scanners := []scanner{{bs: [][]int{}}}
	helper.ReadTestInput(19, func(l string) {
		if l == "" {
			// add scanner
			scanners = append(scanners, scanner{})
			return
		}
		if l[:3] == "---" {
			return
		}
		spl := strings.Split(l, ",")
		bc := []int{}
		for i := range spl {
			c, _ := strconv.Atoi(spl[i])
			bc = append(bc, c)
		}
		scanners[len(scanners)-1].bs = append(scanners[len(scanners)-1].bs, bc)
	})

	i := 0
	for si := range scanners {
		tip := scanners[i]
		_ = tip
		sc := scanners[si]
		for ri := range rotations {
			rot := rotate(sc, rotations[ri])
			_ = rot
			// if checkOverlap(tip, rot) {

			// }
		}
		i++
	}

	return
}

func rotate(sc scanner, r []int) [][]int {
	res := [][]int{}
	for s := range sc.bs {
		c := sc.bs[s]
		res = append(res, []int{
			c[helper.Abs(r[0])-1] * getSign(r[0]),
			c[helper.Abs(r[1])-1] * getSign(r[1]),
			c[helper.Abs(r[2])-1] * getSign(r[2]),
		})
	}
	return res
}
func getSign(i int) int {
	if i < 0 {
		return -1
	}
	return 1
}

func printScanner(sc [][]int) {
	for i := range sc {
		fmt.Println(sc[i])
	}
}
