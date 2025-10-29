package day01

import "fmt"

type Puzzle struct{}

func (p *Puzzle) Part1() string {
    return "Solution to Part 1 of Day 1, 2024"
}

func (p *Puzzle) Part2() string {
    return "Solution to Part 2 of Day 1, 2024"
}

func NewPuzzle() *Puzzle {
    return &Puzzle{}
}