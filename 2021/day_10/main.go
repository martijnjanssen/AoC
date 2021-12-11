package day_10

import (
	"sort"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

var (
	matchings      = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	illegalScores  = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	completeScores = map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

	stack []string
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (int, int) {
	d := 0
	scores := []int{}
	helper.DownloadAndRead(10, func(l string) {
		cs := strings.Split(l, "")
		stack = []string{}
		error := false

	L:
		for _, c := range cs {
			switch c {
			case "{", "(", "[", "<":
				stack = append(stack, c) // append read char to stack
			case "}", ")", "]", ">":
				if matchings[c] == stack[len(stack)-1] { // check top of stack
					stack = stack[0 : len(stack)-1] // pop stack
				} else {
					d += illegalScores[c]
					error = true
					break L
				}
			}
		}
		if error {
			return
		}

		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + completeScores[stack[i]]
		}
		scores = append(scores, score)

	})

	sort.Ints(scores)
	return d, scores[len(scores)/2]
}
