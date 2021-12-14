package day_14

import (
	"math"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

var (
	place map[string]rune
	poly  map[string]int64
	c     map[rune]int64
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (a int, b int) {
	place = map[string]rune{}
	poly = map[string]int64{}
	c = map[rune]int64{}
	polyLoaded := false
	helper.DownloadAndRead(14, func(l string) {
		if l == "" {
			return
		}
		if !polyLoaded {
			polyLoaded = true
			c[rune(l[0])]++
			for i := 0; i < len(l)-1; i++ {
				c[rune(l[i+1])]++
				poly[string(l[i])+string(l[i+1])]++
			}
			return
		}

		ps := strings.Split(l, " -> ")
		place[ps[0]] = rune(ps[1][0])
	})

	for l := range make([]int, 40) {
		newPoly := map[string]int64{}
		for k, v := range poly {
			r := place[k]
			c[r] += v
			newPoly[string(k[0])+string(r)] += v
			newPoly[string(r)+string(k[1])] += v
		}
		poly = newPoly
		if l == 9 {
			a = int(calc(c))
		}
	}
	b = int(calc(c))

	return
}

func calc(c map[rune]int64) int64 {
	var min int64 = math.MaxInt64
	var max int64 = 0
	for _, v := range c {
		min = helper.Min64(min, v)
		max = helper.Max64(max, v)
	}
	return max - min
}
