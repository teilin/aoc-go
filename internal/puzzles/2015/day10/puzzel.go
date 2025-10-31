package day10

import (
	"strconv"
	"strings"
)

type Puzzle struct{}

const PUZZEL_INPUT = "3113322113"

func lookAndSay(s string) string {
	var b strings.Builder
	n := len(s)
	for i := 0; i < n; {
		j := i + 1
		for j < n && s[j] == s[i] {
			j++
		}
		count := j - i
		b.WriteString(strconv.Itoa(count))
		b.WriteByte(s[i])
		i = j
	}
	return b.String()
}

func (p *Puzzle) Part1() string {
	s := PUZZEL_INPUT
	for i := 0; i < 40; i++ {
		s = lookAndSay(s)
	}
	return strconv.Itoa(len(s))
}

func (p *Puzzle) Part2() string {
	s := PUZZEL_INPUT
	for i := 0; i < 50; i++ {
		s = lookAndSay(s)
	}
	return strconv.Itoa(len(s))
}

func New() *Puzzle {
	return &Puzzle{}
}
