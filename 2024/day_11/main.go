package day_11

import (
	"bufio"
	"math"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	stones := map[int]int{}
	helper.ReadLines(buf, func(l string) {
		ls := input.SplitToInt(l, " ")
		for _, v := range ls {
			_, ok := stones[v]
			if !ok {
				stones[v] = 0
			}
			stones[v]++
		}
	})

	for i := 0; i < 75; i++ {
		newStones := make(map[int]int, len(stones))
		for value, amount := range stones {
			numLen := int((math.Floor(math.Log10(float64(value)))) + 1)
			switch {
			case value == 0:
				newStones[1] += amount
			case numLen%2 == 0:
				half := int(math.Pow(10, float64(numLen)/2))
				newStones[value%half] += amount
				newStones[value/half] += amount
			default:
				newStones[value*2024] += amount
			}
		}
		stones = newStones

		if i == 24 {
			a = countStones(stones)
		}
	}

	b = countStones(stones)

	return
}

func countStones(stones map[int]int) int {
	acc := 0
	for _, v := range stones {
		acc += v
	}
	return acc
}
