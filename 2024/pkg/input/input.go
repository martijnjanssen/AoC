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

func SplitToSingleDigitInt(s string) []int {
	ints := make([]int, len(s))
	for i := range ints {
		ints[i] = int(rune(s[i]) - '0')
	}

	return ints
}

func SplitToRune(s string) []rune {
	res := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, rune(s[i]))
	}
	return res
}
