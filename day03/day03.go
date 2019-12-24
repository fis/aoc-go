// Package day03 solves AoC 2019 day 3.
package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc2019-go/util"
)

func Solve(path string) ([]string, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	if len(lines) != 2 {
		return nil, fmt.Errorf("unexpected amount of lines: %v", lines)
	}
	p1, p2 := compute(lines[0], lines[1])
	return []string{strconv.Itoa(p1), strconv.Itoa(p2)}, nil
}

func compute(w1, w2 string) (closest, best int) {
	m := make(map[util.P]int)
	walk(w1, func(x, y, s1 int) {
		p := util.P{x, y}
		if _, ok := m[p]; ok {
			return
		}
		m[p] = s1
	})
	closest, best = -1, -1
	walk(w2, func(x, y, s2 int) {
		s1, ok := m[util.P{x, y}]
		if !ok {
			return
		}
		td := abs(x) + abs(y)
		if closest == -1 || td < closest {
			closest = td
		}
		ts := s1 + s2
		if best == -1 || ts < best {
			best = ts
		}
	})
	return closest, best
}

var dir = map[byte]util.P{
	'L': util.P{-1, 0},
	'R': util.P{1, 0},
	'U': util.P{0, -1},
	'D': util.P{0, 1},
}

func walk(w string, cb func(x, y, s int)) {
	x, y, s := 0, 0, 0
	for _, op := range strings.Split(w, ",") {
		d := dir[op[0]]
		n, _ := strconv.Atoi(op[1:])
		for i := 0; i < n; i++ {
			x, y, s = x+d.X, y+d.Y, s+1
			cb(x, y, s)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
