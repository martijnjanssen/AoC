package day_5

import (
	"bufio"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	cs := [][]int{}

	readConstraints := true
	helper.ReadLines(buf, func(l string) {
		switch {
		case l == "":
			readConstraints = false
			return

		case readConstraints:
			cs = append(cs, input.SplitToInt(l, "|"))
			return

		default:
			ls := input.SplitToInt(l, ",")
			if isValid(cs, ls) {
				a += getMiddleValue(ls)
			} else {
				for !isValid(cs, ls) {
					ls = fix(cs, ls)
				}
				b += getMiddleValue(ls)
			}
		}
	})

	return
}

func isValid(cs [][]int, ls []int) bool {
	for i := range cs {
		a, b := findIndices(ls, cs[i][0], cs[i][1])
		if a == -1 || b == -1 {
			continue
		}
		if a > b {
			return false
		}
	}
	return true
}

func fix(cs [][]int, ls []int) []int {
	for i := range cs {
		a, b := findIndices(ls, cs[i][0], cs[i][1])
		if a == -1 || b == -1 {
			continue
		}
		if a > b {
			ls = append(ls[:a], ls[a+1:]...) // Remove
			ls = append(ls[:b+1], ls[b:]...) // Make room
			ls[b] = cs[i][0]                 // Insert
		}
	}
	return ls
}

func findIndices(ls []int, a int, b int) (int, int) {
	aInd := -1
	bInd := -1
	i := 0
	for i < len(ls) && (aInd == -1 || bInd == -1) {
		if a == ls[i] {
			aInd = i
		}
		if b == ls[i] {
			bInd = i
		}
		i++
	}
	return aInd, bInd
}

func getMiddleValue(ls []int) int {
	return ls[len(ls)/2]
}
