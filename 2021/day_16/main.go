package day_16

import (
	"math"
	"strconv"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

func read(data *string, bits int) string {
	rd := (*data)[:bits]
	(*data) = (*data)[bits:]
	return rd
}

var (
	versionSum int
	hexToBin   = map[byte]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
)

func (r *run) Run() (a int, b int) {
	var data string
	helper.DownloadAndRead(16, func(l string) {
		d := make([]byte, len(l)*4)
		for i := range l {
			copy(d[i*4:], hexToBin[l[i]])
		}
		data = string(d)
	})

	return versionSum, parse(&data)
}

func parse(data *string) int {
	// first 6 header -> 3 version, 3 typeID
	ver, _ := strconv.ParseInt(read(data, 3), 2, 64)
	typ, _ := strconv.ParseInt(read(data, 3), 2, 64)
	versionSum += int(ver)

	switch typ {
	case 4:
		binStr := ""
		for done := false; !done; {
			sig := read(data, 1)
			switch sig {
			case "1": // 4-bit size
				binStr += read(data, 4)
			case "0": // last 4-bit
				binStr += read(data, 4)
				done = true
			}
		}
		num, _ := strconv.ParseInt(binStr, 2, 64)
		return int(num)

	default:
		rs := []int{}
		sig := read(data, 1) // 1 signal bit
		switch sig {
		case "0": // 15-bit size
			siz, _ := strconv.ParseInt(read(data, 15), 2, 64)
			nData := (*data)[:siz]
			*data = (*data)[siz:]
			for len(nData) > 0 {
				rs = append(rs, parse(&nData))
			}
		case "1": // 11-bit size
			siz, _ := strconv.ParseInt(read(data, 11), 2, 64)
			for range make([]int, siz) {
				rs = append(rs, parse(data))
			}
		}
		return calc(int(typ), rs)
	}
}

func calc(typ int, rs []int) int {
	switch typ {
	case 0:
		r := 0
		for i := range rs {
			r += rs[i]
		}
		return r
	case 1:
		r := 1
		for i := range rs {
			r *= rs[i]
		}
		return r
	case 2:
		r := math.MaxInt32
		for i := range rs {
			r = helper.Min(r, rs[i])
		}
		return r
	case 3:
		r := 0
		for i := range rs {
			r = helper.Max(r, rs[i])
		}
		return r
	case 5:
		if rs[0] > rs[1] {
			return 1
		}
		return 0
	case 6:
		if rs[0] < rs[1] {
			return 1
		}
		return 0
	case 7:
		if rs[0] == rs[1] {
			return 1
		}
		return 0
	}
	return 0
}
