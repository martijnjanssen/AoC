package main

import (
	"fmt"
	"strconv"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

// Middle of the window doesn't have to be considered,
// the middle of the window is the same for both instances.
func main() {
	defer helper.Time()()

	ds := []int{}
	helper.DownloadAndRead(1, func(l string) {
		v, _ := strconv.Atoi(l)
		ds = append(ds, v)
	})

	inc := 0
	for i := 1; i < len(ds)-2; i++ {
		if ds[i-1] < ds[i+2] {
			inc++
		}
	}

	fmt.Printf("Increased is: %d\n", inc)
}
