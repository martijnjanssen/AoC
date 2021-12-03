package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

func main() {
	vs := [][]int{}
	size := -1
	helper.DownloadAndRead(3, func(l string) {
		ls := strings.Split(l, "")
		size = len(ls)
		is := []int{}
		for _, l := range ls {
			i, _ := strconv.Atoi(l)
			is = append(is, i)
		}

		vs = append(vs, is)
	})

	first(vs, size)
	second(vs, size)
}

func first(vs [][]int, size int) {
	r := make([]int, size)
	for _, v := range vs {
		for i := 0; i < len(r); i++ {
			if v[i] == 1 {
				r[i] += 1
			} else {
				r[i] -= 1
			}
		}
	}

	for i := range r {
		if r[i] > 0 {
			r[i] = 1
		} else {
			r[i] = 0
		}
	}

	rString := binaryArrayToString(r)
	iString := strings.ReplaceAll(rString, "1", "2")
	iString = strings.ReplaceAll(iString, "0", "1")
	iString = strings.ReplaceAll(iString, "2", "0")

	fmt.Printf("Answer is: %d\n", binaryStringToInt(rString)*binaryStringToInt(iString))
}

func second(vs [][]int, size int) {
	cs := filter(vs, 0, func(c int) bool { return c >= 0 })
	us := filter(vs, 0, func(c int) bool { return c < 0 })

	c := binaryArrayToString(cs[0])
	u := binaryArrayToString(us[0])

	fmt.Printf("Answer is: %d\n", binaryStringToInt(c)*binaryStringToInt(u))
}

func filter(vs [][]int, index int, compFn func(int) bool) [][]int {
	if len(vs) == 1 {
		return vs
	}

	r := [][]int{}
	c := 0
	for _, v := range vs {
		if v[index] == 1 {
			c += 1
		} else {
			c -= 1
		}
	}
	if compFn(c) {
		c = 1
	} else {
		c = 0
	}

	for _, v := range vs {
		if v[index] == c {
			r = append(r, v)
		}
	}

	return filter(r, index + 1, compFn)
}

func binaryArrayToString(bs []int) string {
	bString := ""
	for _, b := range bs {
		bString = fmt.Sprintf("%s%d", bString, b)
	}

	return bString
}

func binaryStringToInt(b string) int {
	r, _ := strconv.ParseInt(b, 2, 64)
	return int(r)
}
