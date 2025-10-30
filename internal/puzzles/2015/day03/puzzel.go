package day03

import (
	"io"
	"os"
	"strconv"
)

type Puzzle struct{}

type Pos struct {
	X int
	Y int
}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day03/input.txt")
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
		return ""
	}
	position := Pos{X: 0, Y: 0}
	visitedMap := make(map[Pos]int)
	visitedMap[position] = 1
	for _, b := range content {
		switch b {
		case '^':
			position.Y++
		case 'v':
			position.Y--
		case '>':
			position.X++
		case '<':
			position.X--
		}
		visitedMap[position]++
	}
	return strconv.Itoa(len(visitedMap))
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	visitedMap := make(map[Pos]int)
	visitedMap[Pos{X: 0, Y: 0}] = 2
	santas := [2]Pos{{X: 0, Y: 0}, {X: 0, Y: 0}}
	for i, b := range content {
		switch b {
		case '^':
			santas[i%2].Y++
		case 'v':
			santas[i%2].Y--
		case '>':
			santas[i%2].X++
		case '<':
			santas[i%2].X--
		}
		visitedMap[santas[i%2]]++
	}
	return strconv.Itoa(len(visitedMap))
}

func New() *Puzzle {
	return &Puzzle{}
}
