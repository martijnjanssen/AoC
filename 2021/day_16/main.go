package day_16

import (
	"fmt"
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
)

func (r *run) Run() (a int, b int) {
	var data string
	helper.DownloadAndRead(16, func(l string) {
		raw := l
		for i := range raw {
			c, _ := strconv.ParseInt(string(raw[i]), 16, 64)
			data += fmt.Sprintf("%04b", c)
		}
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
		r := math.MaxInt
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
