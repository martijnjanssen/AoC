package day_2

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (int, int) {
	hrzA, dthA := 0, 0
	hrzB, aim, dthB := 0, 0, 0

	helper.DownloadAndRead(2, func(l string) {
		s := strings.Split(l, " ")
		a, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			hrzA += a
			hrzB += a
			dthB += aim * a
		case "up":
			dthA -= a
			aim -= a
		case "down":
			dthA += a
			aim += a
		}
	})

	return hrzA * dthA, hrzB * dthB
}
