package day_9

import (
	"bufio"
	_ "embed"

	"github.com/martijnjanssen/aoc/2024/pkg/helper"
	"github.com/martijnjanssen/aoc/2024/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func (r *run) Run(buf *bufio.Reader) (a int, b int) {
	helper.ReadLines(buf, func(l string) {
		a = calculateA(l)
		b = calculateB(l)
	})

	return
}

func toInt(r byte) int {
	return int(r - '0')
}

func calculateA(fs string) int {
	res := 0

	isProcessed := make([]bool, len(fs))

	rest := []int{}
	position := 0
	lastFileIndex := len(fs) - 1
	for i := 0; i < len(fs); i++ {
		if isProcessed[i] {
			continue
		}

		if i%2 == 1 {
			for j := 0; j < toInt(fs[i]); j++ {
				if len(rest) == 0 {
					if isProcessed[lastFileIndex] {
						return res
					}
					lastFileNum := lastFileIndex / 2
					for k := 0; k < toInt(fs[lastFileIndex]); k++ {
						rest = append(rest, lastFileNum)
					}
					isProcessed[lastFileIndex] = true
					lastFileIndex += -2
				}
				res += rest[0] * position
				position++
				rest = rest[1:]
			}
		} else {
			fileNum := i / 2
			for j := 0; j < toInt(fs[i]); j++ {
				res += fileNum * position
				position++
			}
			isProcessed[i] = true
		}
	}

	return res
}

func calculateB(fs string) int {
	res := 0
	isProcessed := make([]bool, len(fs))
	position := 0
	for i := 0; i < len(fs); i++ {
		if i%2 == 1 {
			gap := toInt(fs[i])
			for lb := len(fs) - 1; gap > 0 && lb >= 0; lb -= 2 {
				if isProcessed[lb] || toInt(fs[lb]) > gap {
					continue
				}
				fileNum := lb / 2
				for j := 0; j < toInt(fs[lb]); j++ {
					res += fileNum * position
					position++
				}
				isProcessed[lb] = true
				gap -= toInt(fs[lb])
			}
			for j := 0; j < gap; j++ {
				position++
			}
		} else {
			if isProcessed[i] {
				for j := 0; j < toInt(fs[i]); j++ {
					position++
				}
				continue
			}
			fileNum := i / 2
			for j := 0; j < toInt(fs[i]); j++ {
				res += fileNum * position
				position++
			}
			isProcessed[i] = true
		}
	}

	return res
}
