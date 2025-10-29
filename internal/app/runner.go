package app

import (
	"fmt"
	"github.com/yourusername/aoc-go/pkg/puzzle"
)

type Runner struct {
	PuzzleRegistry puzzle.Registry
}

func (r *Runner) Run(year int, day int) {
	p, err := r.PuzzleRegistry.GetPuzzle(year, day)
	if err != nil {
		fmt.Printf("Error retrieving puzzle for year %d, day %d: %v\n", year, day, err)
		return
	}

	part1Result := p.Part1()
	part2Result := p.Part2()

	fmt.Printf("Year: %d, Day: %d\n", year, day)
	fmt.Printf("Part 1: %s\n", part1Result)
	fmt.Printf("Part 2: %s\n", part2Result)
}