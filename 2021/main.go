package main

import (
	"fmt"
	"time"

	"github.com/martijnjanssen/aoc/2021/day_1"
	"github.com/martijnjanssen/aoc/2021/day_10"
	"github.com/martijnjanssen/aoc/2021/day_11"
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
	}

	defer helper.Time()()
	for i, r := range days[1:] {
		t := time.Now()
		a, b := r.Run()
		fmt.Printf("Day %d:\t\t%s\t\t%d\t\t%d\n", i+1, time.Since(t).String(), a, b)
	}
}
