package day12

import (
	"fmt"

	"alexi.ch/aoc/2026/lib"
)

type Day12 struct {
	s1 uint64
	s2 uint64
}

func New() Day12 {
	return Day12{s1: 0, s2: 0}
}

func (d *Day12) Title() string {
	return "Day 12 - XXXXX"
}

func (d *Day12) Setup() {
	// var lines = lib.ReadLines("data/12-test-data.txt")
	var lines = lib.ReadLines("data/12-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day12) SolveProblem1() {
	d.s1 = 0
}

func (d *Day12) SolveProblem2() {
	d.s2 = 0
}

func (d *Day12) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day12) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
