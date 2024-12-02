package day_2

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
	validA := 0
	validB := 0
	helper.DownloadAndRead(2, func(l string) {
		ls := strings.Split(l, " ")

		if checkSequence(ls) {
			validA++
			validB++
		} else {
			for i := 0; i < len(ls); i++ {
				nLs := append([]string{}, ls[:i]...)
				nLs = append(nLs, ls[i+1:]...)
				if checkSequence(nLs) {
					validB++
					return
				}
			}

		}
	})

	a = validA
	b = validB

	return
}

func checkSequence(ls []string) bool {
	aProbe, _ := strconv.Atoi(ls[0])
	bProbe, _ := strconv.Atoi(ls[1])

	// Set increase or decrease
	asc := aProbe < bProbe

	for i := 0; i < len(ls)-1; i++ {
		a, _ := strconv.Atoi(ls[i])
		b, _ := strconv.Atoi(ls[i+1])

		diff := 0
		if asc {
			diff = b - a
		} else {
			diff = a - b
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
