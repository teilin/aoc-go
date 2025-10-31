package day08

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day08/input.txt")
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
	totalCode := 0
	totalMem := 0

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		totalCode += len(line)

		mem := 0
		n := len(line)
		if n < 2 {
			continue
		}
		for i := 1; i < n-1; {
			if line[i] == '\\' && i+1 < n-1 {
				switch line[i+1] {
				case '\\', '"':
					mem++
					i += 2
				case 'x':
					if i+3 < n-1 {
						mem++
						i += 4 // \ x H H
					} else {
						mem++
						i += 2
					}
				default:
					mem++
					i += 2
				}
			} else {
				mem++
				i++
			}
		}
		totalMem += mem
	}

	return strconv.Itoa(totalCode - totalMem)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	totalCode := 0
	totalEncoded := 0

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		totalCode += len(line)

		enc := 2
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '\\', '"':
				enc += 2
			default:
				enc += 1
			}
		}
		totalEncoded += enc
	}

	return strconv.Itoa(totalEncoded - totalCode)
}

func New() *Puzzle {
	return &Puzzle{}
}
