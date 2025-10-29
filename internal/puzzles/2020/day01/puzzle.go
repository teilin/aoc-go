package puzzles

import "fmt"

type Puzzle struct{}

func (p *Puzzle) Part1() string {
    return "Solution to Part 1 of 2020 Day 1"
}

func (p *Puzzle) Part2() string {
    return "Solution to Part 2 of 2020 Day 1"
}

func NewPuzzle() *Puzzle {
    return &Puzzle{}
}