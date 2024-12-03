package day_4

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type board struct {
	numbers [][]int
	hasWon  bool
}

func (b *board) markNumber(n int) {
	for r := range b.numbers {
		for i := range b.numbers[r] {
			if b.numbers[r][i] == n {
				b.numbers[r][i] = 0
			}
		}
	}
}

func (b *board) checkBingo() bool {
	if b.hasWon {
		return false
	}

	for r := range b.numbers {
		rSum := 0
		cSum := 0
		for i := range b.numbers[r] {
			rSum += b.numbers[r][i]
			cSum += b.numbers[i][r]
		}
		if rSum == 0 || cSum == 0 {
			b.hasWon = true
			return true
		}
	}

	return false
}

func (b *board) getSum() int {
	sum := 0

	for r := range b.numbers {
		for c := range b.numbers[r] {
			sum += b.numbers[r][c]
		}
	}

	return sum
}

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (int, int) {
	numbers := []int{}
	bs := []*board{}
	firstLine := true
	helper.DownloadAndRead(4, func(l string) {
		if firstLine {
			for _, ns := range strings.Split(l, ",") {
				n, _ := strconv.Atoi(ns)
				numbers = append(numbers, n)
			}
			firstLine = false
			return
		}
		if l == "" {
			bs = append(bs, &board{[][]int{}, false})
			return
		}

		ns := []int{}
		is := strings.Fields(l)
		for _, i := range is {
			n, _ := strconv.Atoi(i)
			ns = append(ns, n)
		}

		b := bs[len(bs)-1]
		b.numbers = append(b.numbers, ns)
	})

	return solve(numbers, bs)
}

func solve(numbers []int, bs []*board) (a, b int) {
	bingos := 0
	for _, n := range numbers {
		for _, board := range bs {
			board.markNumber(n)
			if board.checkBingo() {
				if bingos == 0 {
					a = board.getSum() * n
				}
				if bingos == len(bs)-1 {
					b = board.getSum() * n
				}
				bingos++
			}
		}
	}
	return
}
