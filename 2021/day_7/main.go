package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

func main() {
	ps := []int{}
	sum := 0
	helper.DownloadAndRead(7, func(l string) {
		ss := strings.Split(l, ",")
		for _, s := range ss {
			p, _ := strconv.Atoi(s)
			ps = append(ps, p)
			sum += p
		}
	})
	sort.Ints(ps)

	pos := ps[len(ps)/2]
	fuel := 0
	for _, p := range ps {
		fuel += helper.Abs(helper.Diff(pos, p))
	}
	fmt.Printf("Answer is: %d\n", fuel)

	pos = sum / len(ps)
	fuel = 0
	for _, p := range ps {
		d := helper.Abs(helper.Diff(pos, p))
		fuel += (d * (d + 1)) / 2
	}
	fmt.Printf("Answer is: %d\n", fuel)
}
