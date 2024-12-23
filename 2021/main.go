package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"time"

	"log"

	"github.com/martijnjanssen/aoc/2021/day_1"
	"github.com/martijnjanssen/aoc/2021/day_10"
	"github.com/martijnjanssen/aoc/2021/day_11"
	"github.com/martijnjanssen/aoc/2021/day_12"
	"github.com/martijnjanssen/aoc/2021/day_13"
	"github.com/martijnjanssen/aoc/2021/day_14"
	"github.com/martijnjanssen/aoc/2021/day_15"
	"github.com/martijnjanssen/aoc/2021/day_16"
	"github.com/martijnjanssen/aoc/2021/day_17"
	"github.com/martijnjanssen/aoc/2021/day_18"
	"github.com/martijnjanssen/aoc/2021/day_19"
	"github.com/martijnjanssen/aoc/2021/day_20"
	"github.com/martijnjanssen/aoc/2021/day_21"

	// "github.com/martijnjanssen/aoc/2021/day_22"
	// "github.com/martijnjanssen/aoc/2021/day_23"
	// "github.com/martijnjanssen/aoc/2021/day_24"
	// "github.com/martijnjanssen/aoc/2021/day_25"
	"github.com/martijnjanssen/aoc/2021/day_2"
	"github.com/martijnjanssen/aoc/2021/day_3"
	"github.com/martijnjanssen/aoc/2021/day_4"
	"github.com/martijnjanssen/aoc/2021/day_5"
	"github.com/martijnjanssen/aoc/2021/day_6"
	"github.com/martijnjanssen/aoc/2021/day_7"
	"github.com/martijnjanssen/aoc/2021/day_8"
	"github.com/martijnjanssen/aoc/2021/day_9"
	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

func main() {
	if os.Getenv("PPROF") != "" {
		fmt.Println("pprof enabled")
		f, err := os.Create("test.prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	args := os.Args[1:]
	days := []runner.Runner{
		runner.NoOpRunner(),
		day_1.GetRunner(),
		day_2.GetRunner(),
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
		day_14.GetRunner(),
		day_15.GetRunner(),
		day_16.GetRunner(),
		day_17.GetRunner(),
		day_18.GetRunner(),
		day_19.GetRunner(),
		day_20.GetRunner(),
		day_21.GetRunner(),
		// day_22.GetRunner(),
		// day_23.GetRunner(),
		// day_24.GetRunner(),
		// day_25.GetRunner(),
	}

	if len(args) == 0 {
		day := time.Now().Day()
		defer helper.Time()()
		a, b := days[day].Run()
		fmt.Printf("Solutions are: %d\t%d\n", a, b)
		return
	}

	if args[0] == "loop" {
		day, _ := strconv.Atoi(args[1])
		fmt.Printf("Looping day %d\n", day)
		defer helper.Time()()
		for range make([]int, 100) {
			days[day].Run()
		}
		return
	}

	if day, err := strconv.Atoi(args[0]); err == nil && day < len(days) {
		fmt.Printf("Running day %d\n", day)
		defer helper.Time()()
		a, b := days[day].Run()
		fmt.Printf("Solutions are: %d\t%d\n", a, b)
		return
	}

	if args[0] == "all" {
		defer helper.Time()()
		for i := range days[1:] {
			t := time.Now()
			a, b := days[i+1].Run()
			fmt.Printf("Day %d:\t\t%s\t\t%d\t\t%d\n", i+1, time.Since(t).String(), a, b)
		}
		return
	}

	if args[0] == "bench" {
		defer helper.Time()()
		bench(days)
		return
	}

}

func bench(days []runner.Runner) {
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
		fmt.Printf("Day %d:\t\t%s\t\t%s\t\t%d\t\t%d\n", i+1, e.String(), (avg / 100).String(), a, b)
	}
}
