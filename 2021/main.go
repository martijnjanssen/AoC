package main

import (
	"fmt"
	"time"

	"github.com/martijnjanssen/aoc/2021/day_1"
	"github.com/martijnjanssen/aoc/2021/day_10"
	"github.com/martijnjanssen/aoc/2021/day_11"
	"github.com/martijnjanssen/aoc/2021/day_12"
	"github.com/martijnjanssen/aoc/2021/day_13"
	"github.com/martijnjanssen/aoc/2021/day_3"
	"github.com/martijnjanssen/aoc/2021/day_4"
	"github.com/martijnjanssen/aoc/2021/day_5"
	"github.com/martijnjanssen/aoc/2021/day_6"
	"github.com/martijnjanssen/aoc/2021/day_7"
	"github.com/martijnjanssen/aoc/2021/day_8"
	"github.com/martijnjanssen/aoc/2021/day_9"
	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

func main() {

	days := []runner.Runner{
		runner.NoOpRunner(),
		day_1.GetRunner(),
		runner.NoOpRunner(),
		day_3.GetRunner(),
		day_4.GetRunner(),
		day_5.GetRunner(),
		day_6.GetRunner(),
		day_7.GetRunner(),
		day_8.GetRunner(),
		day_9.GetRunner(),
		day_10.GetRunner(),
		day_11.GetRunner(),
		day_12.GetRunner(),
		day_13.GetRunner(),
	}

	defer helper.Time()()
	for i, r := range days[1:] {
		start := time.Now()
		a, b := r.Run()
		e := time.Since(start)
		avg := e
		for _ = range make([]int, 99) {
			start := time.Now()
			r.Run()
			t := time.Since(start)
			avg += t
			if t < e {
				e = t
			}
		}
		fmt.Printf("Day %d:\t\t%s\t\t%s\t\t%d\t\t%d\n", i+1, e.String(), (avg/100).String(), a, b)
	}
}
