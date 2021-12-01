package main

import (
	"strconv"

	"github.com/martijnjanssen/aoc/pkg/helper"

	"github.com/sirupsen/logrus"
)

func main() {
	ds := []int{}

	helper.DownloadAndRead(1, func(l string) {
		v, err := strconv.Atoi(l)
		if err != nil {
			logrus.Fatalf("Unable to read line: %s", err)
		}

		ds = append(ds, v)
	})

	inc := 0
	for i := 1; i < len(ds)-2; i++ {
		if ds[i-1]+ds[i]+ds[i+1] < ds[i]+ds[i+1]+ds[i+2] {
			inc++
		}
	}

	logrus.Infof("Increased is: %d", inc)
}
