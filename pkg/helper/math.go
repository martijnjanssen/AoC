package helper

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Min64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Diff(a int, b int) int {
	return a - b
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a

	}
	return a
}
