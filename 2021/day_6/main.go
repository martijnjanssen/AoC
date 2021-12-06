package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

func main() {
	fs := make([]int, 9)

	helper.DownloadAndRead(6, func(l string) {
		ds := strings.Split(l, ",")
		for _, d := range ds {
			i, _ := strconv.Atoi(d)
			fs[i]++
		}
	})

	for r := range make([]int, 256) {
		newFs := fs[0]
		for i := 1; i < len(fs); i++ {
			fs[i-1] = fs[i]
		}
		fs[6] += newFs
		fs[8] = newFs

		if r == 79 || r == 255 {
			printResult(fs)
		}
	}

}

func printResult(fs []int) {
	c := 0
	for i := range fs {
		c += fs[i]
	}

	fmt.Printf("Answer is: %d\n", c)

}
