package day_6

import (
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
	fs := make([]int, 9)
	helper.DownloadAndRead(6, func(l string) {
		ds := strings.Split(l, ",")
		for _, d := range ds {
			i, _ := strconv.Atoi(d)
			fs[i]++
		}
	})

	for r := range make([]int, 256) {
		newFs := fs[0]
		for i := 1; i < len(fs); i++ {
			fs[i-1] = fs[i]
		}
		fs[6] += newFs
		fs[8] = newFs

		if r == 79 {
			a = getResult(fs)
		} else if r == 255 {
			b = getResult(fs)
		}
	}

	return
}

func getResult(fs []int) int {
	c := 0
	for i := range fs {
		c += fs[i]
	}
	return c
}
