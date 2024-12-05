package day_5

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	cs := [][]string{}

	readConstraints := true
	helper.ReadLines(buf, func(l string) {
		if l == "" {
			readConstraints = false
			return
		}
		if readConstraints {
			ac, bc, _ := strings.Cut(l, "|")
			cs = append(cs, []string{ac, bc})
		} else {
			if isValid(cs, l) {
				a += getMiddleValue(l)
			} else {
				for !isValid(cs, l) {
					l = fix(cs, l)
				}
				b += getMiddleValue(l)
			}
		}
	})

	return
}

func isValid(cs [][]string, l string) bool {
	for _, c := range cs {
		aIndex := strings.Index(l, c[0])
		bIndex := strings.Index(l, c[1])
		if aIndex == -1 || bIndex == -1 {
			continue
		}
		if aIndex > bIndex {
			return false
		}
	}
	return true
}

func fix(cs [][]string, l string) string {
	for _, c := range cs {
		aIndex := strings.Index(l, c[0])
		bIndex := strings.Index(l, c[1])
		if aIndex == -1 || bIndex == -1 {
			continue
		}
		if aIndex > bIndex {
			l = strings.Replace(l, c[0], "", 1)
			l = strings.Replace(l, c[1], c[0]+","+c[1], 1)
			l = strings.Replace(l, ",,", ",", 1)
			l = strings.Trim(l, ",")
		}
	}
	return l
}

func getMiddleValue(l string) int {
	spl := strings.Split(l, ",")
	res, _ := strconv.Atoi(spl[len(spl)/2])
	return res
}
