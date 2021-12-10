package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

var (
	matchings      = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	illegalScores  = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	completeScores = map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}

	stack []string
)

func main() {
	defer helper.Time()()

	r := 0
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
					r += illegalScores[c]
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

	fmt.Printf("Answer is: %d\n", r)

	sort.Ints(scores)
	fmt.Printf("Answer is: %d\n", scores[len(scores)/2])
}
