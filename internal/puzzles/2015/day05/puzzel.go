package day05

import (
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day05/input.txt")
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
	count := 0
	for _, line := range strings.Split(string(content), "\n") {
		if isNice(string(line)) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	count := 0
	for _, line := range strings.Split(string(content), "\n") {
		if isEvenNicer(string(line)) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func New() *Puzzle {
	return &Puzzle{}
}

func isNice(s string) bool {
	if countVowels(s) < 3 {
		return false
	}
	if countAdjacentSameLetters(s) < 1 {
		return false
	}
	if hasForbiddenSubstring(s) {
		return false
	}
	return true
}

func isEvenNicer(s string) bool {
	if hasRepeatWithOneBetween(s) {
		return false
	}
	if countAdjacentSameLetters(s) < 2 {
		return false
	}
	return true
}

func hasRepeatWithOneBetween(s string) bool {
	runes := []rune(s)
	for i := 0; i+2 < len(runes); i++ {
		if runes[i] == runes[i+2] {
			return true
		}
	}
	return false
}

func countVowels(s string) int {
	count := 0
	for _, r := range s {
		switch unicode.ToLower(r) {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func countAdjacentSameLetters(s string) int {
	count := 0
	var prev rune
	first := true
	for _, r := range s {
		if !first && r == prev {
			count++
		}
		prev = r
		first = false
	}
	return count
}

func hasForbiddenSubstring(s string) bool {
	for _, f := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, f) {
			return true
		}
	}
	return false
}
