package day01

import (
	"fmt"
	"io"
	"os"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day01/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (p *Puzzle) Part1() string {
	content, err := readInput()
	if err != nil {
		return fmt.Sprintf("Error reading input: %v", err)
	}
	floor := 0
	for _, char := range content {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return fmt.Sprintf("Result of Part 1: %d", floor)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return fmt.Sprintf("Error reading input: %v", err)
	}
	floor := 0
	for i, char := range content {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return fmt.Sprintf("Result of Part 2: %d", i+1)
		}
	}
	return "Result of Part 2: Not found"
}

func New() *Puzzle {
	return &Puzzle{}
}
