package day_23

import (
	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (a int, b int) {
	helper.DownloadAndRead(23, func(l string) {

	})

	return
}
