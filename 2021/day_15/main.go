package day_15

import (
	"sort"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var (
	cave     [][]int
	expanded [][]int
	mY       int
	mX       int
)

func (r *run) Run() (a int, b int) {
	cave = [][]int{}
	helper.DownloadAndRead(15, func(l string) {
		spl := strings.Split(l, "")
		ns := []int{}
		for i := range spl {
			n, _ := strconv.Atoi(spl[i])
			ns = append(ns, n)
		}
		cave = append(cave, ns)
	})
	expanded = make([][]int, len(cave)*5)
	for i := range expanded {
		expanded[i] = make([]int, len(cave[0])*5)
	}

	mY = len(cave)
	mX = len(cave[0])

	// go func() {
	// 	for true {
	// 		for y := 0; y < mY; y++ {
	// 			for x := 0; x < mX; x++ {
	// 				fmt.Printf("%d\t", expanded[y][x])
	// 			}
	// 			fmt.Println("")
	// 		}
	// 		fmt.Println("")
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	todo := [][]int{{0, 0, 0}}
	for len(todo) > 0 {
		tip := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		y := tip[0]
		x := tip[1]
		r := tip[2]
		if expanded[y][x] != 0 && r >= expanded[y][x] {
			continue
		}
		expanded[y][x] = r

		rs := getRisks(y, x)
		// fmt.Println(rs)
		for i := 0; i < len(rs); i++ {
			nr := rs[i][2]
			if nr == -1 {
				continue
			}
			ny := rs[i][0]
			nx := rs[i][1]
			todo = insert(todo, []int{ny, nx, r + nr})
		}
	}

	a = expanded[mY-1][mX-1]
	b = expanded[mY*5-1][mX*5-1]

	return
}

func insert(s [][]int, e []int) [][]int {
	i := sort.Search(len(s), func(i int) bool {
		return s[i][2] <= e[2]
	})
	s = append(s, nil)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func getRisks(y int, x int) [][]int {
	rs := [][]int{}
	rs = insert(rs, []int{y, x - 1, onGrid(y, x-1)})
	rs = insert(rs, []int{y - 1, x, onGrid(y-1, x)})
	rs = insert(rs, []int{y, x + 1, onGrid(y, x+1)})
	return insert(rs, []int{y + 1, x, onGrid(y+1, x)})
}

func onGrid(y int, x int) int {
	if y >= 0 && y < len(expanded) && x >= 0 && x < len(expanded[0]) {
		return (cave[y%mY][x%mX]+y/mY+x/mX-1)%9 + 1
	}
	return -1
}
