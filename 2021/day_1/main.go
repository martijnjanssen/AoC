package day_1

import (
	"strconv"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

// Middle of the window doesn't have to be considered,
// the middle of the window is the same for both instances.
type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (a int, b int) {
	ds := []int{}
	helper.DownloadAndRead(1, func(l string) {
		v, _ := strconv.Atoi(l)
		ds = append(ds, v)
	})

	for i := 1; i < len(ds)-1; i++ {
		if ds[i-1] < ds[i+1] {
			a++
		}
	}
	for i := 1; i < len(ds)-2; i++ {
		if ds[i-1] < ds[i+2] {
			b++
		}
	}

	return
}
