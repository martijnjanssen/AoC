package day_18

import (
	"strconv"

	"github.com/martijnjanssen/aoc/pkg/helper"
	"github.com/martijnjanssen/aoc/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

type ls struct {
	p     *ls
	left  *ls
	right *ls

	num int
}

type rank struct {
	mag int
	num string
}

func (r *run) Run() (a int, b int) {
	var ast *ls
	numbers := []string{}
	helper.DownloadAndRead(18, func(l string) {
		numbers = append(numbers, l)
		if ast == nil {
			ast = createAST(l, 0, nil)
			return
		}

		combAST := &ls{left: ast, num: -1}
		combAST.right = createAST(l, 0, combAST)
		for traverse(combAST) {
		}
		ast = combAST
		a = magnitude(ast)
	})

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}

			combAST := createAST("["+numbers[i]+","+numbers[j]+"]", 0, nil)
			for traverse(combAST) {
			}
			b = helper.Max(b, magnitude(combAST))
		}
	}

	return
}

func getString(t *ls) string {
	if t.left != nil {
		return "[" + getString(t.left) + "," + getString(t.right) + "]"
	}

	return strconv.Itoa(t.num)
}

func magnitude(t *ls) int {
	if t.num > -1 {
		return t.num
	}
	return 3*magnitude(t.left) + 2*magnitude(t.right)
}

func createAST(n string, d int, p *ls) *ls {
	c := 0
L:
	for i := 0; i < len(n); i++ {
		switch n[i] {
		case '[':
			c++
		case ']':
			c--
		case ',':
			if c == 1 {
				c = i
				break L
			}
		default:
			if c == 0 {
				nr, _ := strconv.Atoi(n)
				return &ls{num: nr, p: p}
			}
		}
	}
	node := &ls{num: -1, p: p}
	node.left = createAST(n[1:c], d+1, node)
	node.right = createAST(n[c+1:len(n)-1], d+1, node)
	return node
}

func traverse(t *ls) bool {
	if explode(t, 0) {
		return true
	}

	return split(t, 0)
}

func explode(t *ls, d int) bool {
	if t.left != nil { // Still in nodes
		if t.left.num != -1 && t.right.num != -1 && d >= 4 {
			addSide(t, "left", t.left.num)
			addSide(t, "right", t.right.num)
			t.left = nil
			t.right = nil
			t.num = 0
			return true
		}
		return explode(t.left, d+1) || explode(t.right, d+1)
	}

	return false
}

func split(t *ls, d int) bool {
	if t == nil {
		return false
	}
	if split(t.left, d+1) {
		return true
	}
	if split(t.right, d+1) {
		return true
	}

	if t.num > 9 {
		t.left = &ls{num: t.num / 2, p: t}
		if t.num%2 == 0 {
			t.right = &ls{num: t.num / 2, p: t}
		} else {
			t.right = &ls{num: (t.num-1)/2 + 1, p: t}
		}
		t.num = -1
		return true
	}

	return false
}

func addSide(t *ls, s string, n int) {
	var prev *ls
	if s == "left" {
		for t.p != nil {
			prev = t
			t = t.p
			if prev == t.right {
				prev = t
				t = t.left
				for t.num == -1 {
					prev = t
					t = t.right
				}
				t.num += n
				return
			}
		}
	} else {
		for t.p != nil {
			prev = t
			t = t.p
			if prev == t.left {
				prev = t
				t = t.right
				for t.num == -1 {
					prev = t
					t = t.left
				}
				t.num += n
				return
			}
		}
	}
}
