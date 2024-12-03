package day_12

import (
	"strings"

	"github.com/martijnjanssen/aoc/2021/pkg/helper"
	"github.com/martijnjanssen/aoc/2021/pkg/runner"
)

type run struct{}

func GetRunner() runner.Runner {
	return &run{}
}

var (
	caves map[string]*cave
)

type cave struct {
	name    string
	isSmall bool
	next    []*cave
}

type path struct {
	c       *cave
	visited string
	double  string
}

func (c *cave) addNext(c1 *cave) {
	if c.name == "end" || c1.name == "start" {
		return
	}
	c.next = append(c.next, c1)
}

func (r *run) Run() (a int, b int) {
	caves = map[string]*cave{}
	helper.DownloadAndRead(12, func(l string) {
		es := strings.Split(l, "-")
		for _, e := range es {
			if _, ok := caves[e]; !ok {
				caves[e] = &cave{e, e == strings.ToLower(e), []*cave{}}
			}
		}
		caves[es[0]].addNext(caves[es[1]])
		caves[es[1]].addNext(caves[es[0]])
	})

	paths := []*path{{c: caves["start"], visited: ""}}
	for len(paths) > 0 {
		tip := paths[len(paths)-1]
		paths = paths[:len(paths)-1]

		for i := range tip.c.next {
			if len(tip.c.next[i].next) == 0 {
				if tip.double == "" {
					a++
				}
				b++
				continue
			}
			p := tip.visitCave(tip.c.next[i])
			if p != nil {
				paths = append(paths, p)
			}
		}
	}

	return
}

func (p *path) visitCave(c *cave) *path {
	if !c.isSmall {
		return &path{c, p.visited, p.double} // Allow big caves
	}

	if !strings.Contains(p.visited, c.name) {
		return &path{c, p.visited + "," + c.name, p.double} // Allow small non-visited caves
	}

	if p.double == "" {
		return &path{c, p.visited, c.name} // Allow double visit for one small cave
	}

	return nil
}
