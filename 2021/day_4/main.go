package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
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

func main() {
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

	solve(numbers, bs)
}

func solve(numbers []int, bs []*board) {
	bingos := 0
	for _, n := range numbers {
		for _, b := range bs {
			b.markNumber(n)
			if b.checkBingo() {
				if bingos == 0 {
					fmt.Printf("First bingo: %d\n", b.getSum()*n)
				}
				if bingos == len(bs)-1 {
					fmt.Printf("Last bingo: %d\n", b.getSum()*n)
				}
				bingos++
			}
		}
	}
}
