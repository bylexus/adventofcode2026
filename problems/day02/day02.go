package day02

import (
	"fmt"

	"alexi.ch/aoc/2026/lib"
)

type Day02 struct {
	s1 uint64
	s2 uint64
}

func New() Day02 {
	return Day02{s1: 0, s2: 0}
}

func (d *Day02) Title() string {
	return "Day 02 - XXXXX"
}

func (d *Day02) Setup() {
	// var lines = lib.ReadLines("data/02-test-data.txt")
	var lines = lib.ReadLines("data/02-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day02) SolveProblem1() {
	d.s1 = 0
}

func (d *Day02) SolveProblem2() {
	d.s2 = 0
}

func (d *Day02) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day02) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
