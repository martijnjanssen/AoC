package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martijnjanssen/aoc/pkg/helper"
)

type direction string

var (
	FORWARD direction = "forward"
	DOWN    direction = "down"
	UP      direction = "up"
)

var (
	horizontal = 0
	aim        = 0
	depth      = 0
)

type action struct {
	command direction
	amount  int
}

func main() {
	defer helper.Time()()

	as := []action{}
	helper.DownloadAndRead(2, func(l string) {
		s := strings.Split(l, " ")
		a, _ := strconv.Atoi(s[1])
		as = append(as, action{direction(s[0]), a})
	})

	for _, a := range as {
		switch a.command {
		case FORWARD:
			horizontal += a.amount
			depth += aim * a.amount
			break
		case UP:
			aim -= a.amount
			break
		case DOWN:
			aim += a.amount
			break
		}
	}

	fmt.Printf("Multiplication is: %d\n", horizontal*depth)
}
