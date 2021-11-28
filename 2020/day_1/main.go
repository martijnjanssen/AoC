package main

import (
	"strconv"

	"github.com/martijnjanssen/aoc/pkg/input"
	log "github.com/sirupsen/logrus"
)

func main() {
	goal := 2020
	items := []int{}

	r := input.OpenReader("input.txt")
	input.ReadLines(r, func(l string) {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("line was not an integer: %s\n%s", err, l)
		}
		if n > goal {
			log.Warnf("value was over %d: %d", goal, n)
			return
		}

		items = append(items, n)
	})

	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items)-1; j++ {
			for k := j + 1; k < len(items)-1; k++ {
				a := items[i]
				b := items[j]
				c := items[k]
				if a+b+c == goal {
					log.Infof("Answer found: %d", a*b*c)
					return
				}
			}
		}
	}
}
