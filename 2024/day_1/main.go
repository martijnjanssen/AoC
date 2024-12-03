package day_1

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	ls := []int{}
	rs := []int{}

	rMap := map[int]int{}
	helper.ReadLines(buf, func(l string) {
		a, b, ok := strings.Cut(l, "   ")
		if !ok {
			panic("wrong format")
		}
		lv, _ := strconv.Atoi(a)
		ls = append(ls, lv)
		rv, _ := strconv.Atoi(b)
		rs = append(rs, rv)

		v := rMap[rv]
		rMap[rv] = v + 1
	})

	slices.Sort(ls)
	slices.Sort(rs)

	aSum := 0
	for i := 0; i < len(ls); i++ {
		res := ls[i] - rs[i]
		if res < 0 {
			res *= -1
		}
		aSum += res
	}
	a = aSum

	bSum := 0
	for _, lv := range ls {
		bSum += lv * rMap[lv]
	}
	b = bSum

	return
}
