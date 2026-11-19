package day09

import (
	"fmt"

	"alexi.ch/aoc/2026/lib"
)

type Day09 struct {
	s1 uint64
	s2 uint64
}

func New() Day09 {
	return Day09{s1: 0, s2: 0}
}

func (d *Day09) Title() string {
	return "Day 09 - XXXXX"
}

func (d *Day09) Setup() {
	// var lines = lib.ReadLines("data/09-test-data.txt")
	var lines = lib.ReadLines("data/09-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day09) SolveProblem1() {
	d.s1 = 0
}

func (d *Day09) SolveProblem2() {
	d.s2 = 0
}

func (d *Day09) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day09) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
