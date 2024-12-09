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
	helper.ReadLines(buf, func(l string) {
		fs := input.SplitToInt(l, "")
		a = calculateA(fs)
		b = calculateB(fs)
	})

	return
}

func calculateA(fs []int) int {
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
			// If it is an empty space, put an item in the rest buffer,
			// from which we take file parts to put in the empty spaces
			// as early as possible. If we cannot put the entire rest
			// buffer into the empty space, the 'rest' part is  perserved
			// and we will get to it in the next empty space.
			for j := 0; j < fs[i]; j++ {
				if isProcessed[lastFileIndex] {
					return res
				}
				if len(rest) == 0 {
					lastFileNum := lastFileIndex / 2
					for k := 0; k < fs[lastFileIndex]; k++ {
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
			// In case of a file, we write it as usual
			fileNum := i / 2
			for j := 0; j < fs[i]; j++ {
				res += fileNum * position
				position++
			}
			isProcessed[i] = true
		}
	}

	return res
}

func calculateB(fs []int) int {
	res := 0
	isProcessed := make([]bool, len(fs))
	position := 0
	for i := 0; i < len(fs); i++ {
		// For an empty space, we start checking from the back of
		// the list of files which file would fit in that empty space.
		if i%2 == 1 {
			gap := fs[i]
			// While we have a gap > 0 and we can still use an earlier file
			for lb := len(fs) - 1; gap > 0 && lb >= 0; lb -= 2 {
				// If the file is already processed or it is too large
				if isProcessed[lb] || fs[lb] > gap {
					continue
				}
				fileNum := lb / 2
				for j := 0; j < fs[lb]; j++ {
					res += fileNum * position
					position++
				}
				isProcessed[lb] = true
				gap -= fs[lb]
			}
			// For the remaining gap, write empty space
			for j := 0; j < gap; j++ {
				position++
			}
		} else {
			if isProcessed[i] {
				// If the file is already processed, it means that it
				// was used to fill an empty space before, so we'll
				// write an empty space instead of the actual file.
				for j := 0; j < fs[i]; j++ {
					position++
				}
				continue
			}

			// Normal operation for a file that is not yet processed.
			fileNum := i / 2
			for j := 0; j < fs[i]; j++ {
				res += fileNum * position
				position++
			}
			isProcessed[i] = true
		}
	}

	return res
}
