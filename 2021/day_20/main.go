package day_20

import (
	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var (
	flashing    bool
	outbounds   = 0
	enhancement []int
	image       [][]int
	mul         = []int{256, 128, 64, 32, 16, 8, 4, 2, 1}
)

func (r *run) Run() (a int, b int) {
	enhancement = []int{}
	image = [][]int{}
	helper.DownloadAndRead(20, func(l string) {
		if l == "" {
			return
		}
		if len(enhancement) == 0 {
			for i := range l {
				if l[i] == '#' {
					enhancement = append(enhancement, 1)
				} else {
					enhancement = append(enhancement, 0)
				}
			}
			if enhancement[0] == 1 && enhancement[len(enhancement)-1] == 0 {
				flashing = true
			}
			return
		}

		imageLine := []int{}
		for i := range l {
			if l[i] == '#' {
				imageLine = append(imageLine, 1)
			} else {
				imageLine = append(imageLine, 0)
			}
		}
		image = append(image, imageLine)
	})

	for i := range make([]int, 50) {
		if flashing { // out of bounds points should be 1 when the infinity side is flashing
			outbounds = i % 2
		}
		newImage := make([][]int, len(image)+2)
		for y := range newImage {
			newImage[y] = make([]int, len(image[0])+2)
			for x := range newImage[y] {
				newImage[y][x] = enhance(y-1, x-1)
			}
		}
		image = newImage
		if i == 1 {
			a = lit(image)
		}
	}
	b = lit(image)

	return
}

func lit(image [][]int) int {
	acc := 0
	for y := range image {
		for x := range image[y] {
			acc += image[y][x]
		}
	}
	return acc
}

func enhance(y int, x int) int {
	return enhancement[getPixel(y-1, x-1)*256+getPixel(y-1, x)*128+getPixel(y-1, x+1)*64+
		getPixel(y, x-1)*32+getPixel(y, x)*16+getPixel(y, x+1)*8+
		getPixel(y+1, x-1)*4+getPixel(y+1, x)*2+getPixel(y+1, x+1)]
}

func getPixel(y int, x int) int {
	if y < 0 || y >= len(image) || x < 0 || x >= len(image[y]) {
		return outbounds
	}
	return image[y][x]
}
