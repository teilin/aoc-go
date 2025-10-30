package day02

import (
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day02/input.txt")
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
	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		box := strings.Split(line, "x")
		l, _ := strconv.Atoi(box[0])
		w, _ := strconv.Atoi(box[1])
		h, _ := strconv.Atoi(box[2])
		area, extra := surfaceArea(l, w, h)
		sum += area + extra
	}
	return strconv.Itoa(sum)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		box := strings.Split(line, "x")
		l, _ := strconv.Atoi(box[0])
		w, _ := strconv.Atoi(box[1])
		h, _ := strconv.Atoi(box[2])
		smallestDistance := 2 * slices.Min([]int{l + w, w + h, h + l})
		smallestPerimeter := slices.Min([]int{2*l + 2*w, 2*w + 2*h, 2*h + 2*l})
		smallest := slices.Min([]int{smallestDistance, smallestPerimeter})
		volume := l * w * h
		sum += smallest + volume
	}
	return strconv.Itoa(sum)
}

func New() *Puzzle {
	return &Puzzle{}
}

func surfaceArea(l, w, h int) (int, int) {
	sides := []int{l * w, w * h, h * l}
	return 2*l*w + 2*w*h + 2*h*l, slices.Min(sides)
}
