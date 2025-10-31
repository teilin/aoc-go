package day11

import "strings"

type Puzzle struct{}

const PUZZEL_INPUT = "vzbxkghb"

func (p *Puzzle) Part1() string {
	return nextValidPassword(PUZZEL_INPUT)
}

func (p *Puzzle) Part2() string {
	return nextValidPassword(nextValidPassword(PUZZEL_INPUT))
}

func New() *Puzzle {
	return &Puzzle{}
}

func incrementPassword(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		b[i]++
		if b[i] > 'z' {
			b[i] = 'a'
			continue
		}
		// If we hit a disallowed letter, bump it and set following chars to 'a'
		if b[i] == 'i' || b[i] == 'o' || b[i] == 'l' {
			b[i]++
		}
		for j := i + 1; j < len(b); j++ {
			b[j] = 'a'
		}
		break
	}
	return string(b)
}

func hasStraight(s string) bool {
	b := []byte(s)
	for i := 0; i+2 < len(b); i++ {
		if b[i+1] == b[i]+1 && b[i+2] == b[i]+2 {
			return true
		}
	}
	return false
}

func hasTwoDifferentPairs(s string) bool {
	b := []byte(s)
	pairs := make(map[byte]struct{})
	for i := 0; i+1 < len(b); {
		if b[i] == b[i+1] {
			pairs[b[i]] = struct{}{}
			i += 2 // skip to avoid overlapping
		} else {
			i++
		}
	}
	return len(pairs) >= 2
}

func containsInvalid(s string) bool {
	return strings.ContainsAny(s, "iol")
}

func isValidPassword(s string) bool {
	if containsInvalid(s) {
		return false
	}
	if !hasStraight(s) {
		return false
	}
	if !hasTwoDifferentPairs(s) {
		return false
	}
	return true
}

func nextValidPassword(start string) string {
	p := incrementPassword(start)
	for !isValidPassword(p) {
		p = incrementPassword(p)
	}
	return p
}
