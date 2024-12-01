package day_1

import (
	"slices"
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
	ls := []int{}
	rs := []int{}

	rMap := map[int]int{}
	helper.DownloadAndRead(1, func(l string) {
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
