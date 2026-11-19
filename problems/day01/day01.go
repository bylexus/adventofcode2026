package day01

import (
	"fmt"

	"alexi.ch/aoc/2026/lib"
)

type Day01 struct {
	s1 int
	s2 int
}

func New() Day01 {
	return Day01{}
}

func (d *Day01) Title() string {
	return "Day 01 - XXXXXXX"
}

func (d *Day01) Setup() {
	var lines = lib.ReadLines("data/01-data.txt")
	// var lines = lib.ReadLines("data/01-test-data.txt")
	for _, line := range lines {
		line = line
		fmt.Println("Here comes some parsing ...")
	}
}

func (d *Day01) SolveProblem1() {
	d.s1 = 0
}

func (d *Day01) SolveProblem2() {
	d.s2 = 0
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
