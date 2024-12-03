package day_8

import (
	"math"
	"sort"
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

var numbers map[string]int
var rNumbers map[int]string

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run() (int, int) {
	occurrences := 0
	acc := 0
	helper.DownloadAndRead(8, func(l string) {
		numbers = map[string]int{}
		rNumbers = map[int]string{}
		lSplit := strings.Split(l, "|")
		iSplit := strings.Split(strings.TrimSpace(lSplit[0]), " ")
		oSplit := strings.Split(strings.TrimSpace(lSplit[1]), " ")
		// Sort such that the 6 comes before the 5,
		// length 6 strings should be figured out first.
		sort.Slice(iSplit, func(i, j int) bool {
			if len(iSplit[i]) == 5 {
				return false
			}
			if len(iSplit[j]) == 5 {
				return true
			}
			return len(iSplit[i]) < len(iSplit[j])
		})

		for _, s := range iSplit {
			s = sortString(s)
			switch len(s) {
			case 2:
				storeNums(1, s)
			case 3:
				storeNums(7, s)
			case 4:
				storeNums(4, s)
			case 7:
				storeNums(8, s)
			case 6:
				if hasActivatedSegments(s, rNumbers[4]) {
					storeNums(9, s)
				} else if hasActivatedSegments(s, rNumbers[1]) {
					storeNums(0, s)
				} else {
					storeNums(6, s)
				}
			case 5:
				if hasActivatedSegments(s, rNumbers[1]) {
					storeNums(3, s)
				} else if hasActivatedSegments(rNumbers[6], s) {
					storeNums(5, s)
				} else {
					storeNums(2, s)
				}
			}
		}

		for i, s := range oSplit {
			// Calculation for part 1
			if len(s) == 7 || len(s) == 4 || len(s) == 3 || len(s) == 2 {
				occurrences++
			}

			// Results part 2
			acc += numbers[sortString(oSplit[i])] * int(math.Pow10(3-i))
		}
	})

	return occurrences, acc
}

func storeNums(n int, s string) {
	numbers[s] = n
	rNumbers[n] = s
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func hasActivatedSegments(c string, h string) bool {
	for _, l := range h {
		if strings.IndexRune(c, l) == -1 {
			return false
		}
	}
	return true
}
