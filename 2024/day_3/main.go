package day_3

import (
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

const doCmd = "do()"
const dontCmd = "don't()"
const mulStartCmd = "mul("

func (r *run) Run() (a int, b int) {
	accA := 0
	accB := 0
	do := true
	helper.DownloadAndRead(3, func(l string) {
		for len(l) > 0 {
			switch {
			case len(dontCmd) <= len(l) && l[:len(dontCmd)] == dontCmd:
				do = false
				l = l[len(dontCmd):]

			case len(doCmd) <= len(l) && l[:len(doCmd)] == doCmd:
				do = true
				l = l[len(doCmd):]

			case len(mulStartCmd) <= len(l) && l[:len(mulStartCmd)] == mulStartCmd:
				a, aRest, _ := strings.Cut(l[len(mulStartCmd):], ",")
				iA, err := strconv.Atoi(a)
				if err != nil {
					l = l[len(mulStartCmd):]
					continue
				}
				b, bRest, _ := strings.Cut(aRest, ")")
				iB, err := strconv.Atoi(b)
				if err != nil {
					l = l[len(mulStartCmd):]
					continue
				}

				accA += iA * iB
				if do {
					accB += iA * iB
				}

				l = bRest
				continue

			default:
				l = l[1:]
			}

		}
	})

	a = accA
	b = accB

	return
}
