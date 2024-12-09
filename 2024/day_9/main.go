package day_9

import (
	"bufio"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/input"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	rest := []int{}
	helper.ReadLines(buf, func(l string) {
		fs := input.SplitToInt(l, "")

		position := 0
		lastFileIndex := len(fs) - 1
		for i := 0; i < len(fs); i++ {
			if fs[i] == -1 {
				continue
			}

			if i%2 == 1 {
				for j := 0; j < fs[i]; j++ {
					if len(rest) == 0 {
						if fs[lastFileIndex] == -1 {
							return
						}
						lastFileNum := lastFileIndex / 2
						for k := 0; k < fs[lastFileIndex]; k++ {
							rest = append(rest, lastFileNum)
						}
						fs[lastFileIndex] = -1
						lastFileIndex += -2
					}
					a += rest[0] * position
					position++
					rest = rest[1:]
				}
			} else {
				fileNum := i / 2
				for j := 0; j < fs[i]; j++ {
					a += fileNum * position
					position++
				}
				fs[i] = -1
			}
		}
	})

	return
}
