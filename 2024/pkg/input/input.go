package input

import (
	"strconv"
	"strings"
)

func SplitToInt(s string, sep string) []int {
	strs := strings.Split(s, sep)
	ints := make([]int, len(strs))
	var err error
	for i := range ints {
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			panic(err)
		}
	}

	return ints
}
