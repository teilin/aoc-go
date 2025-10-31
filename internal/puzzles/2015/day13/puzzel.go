package day13

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day13/input.txt")
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
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	hap := make(map[string]map[string]int)
	namesMap := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Example line:
		// Alice would gain 54 happiness units by sitting next to Bob.
		toks := strings.Fields(line)
		if len(toks) < 11 {
			continue
		}
		name := toks[0]
		gainOrLose := toks[2]
		valueStr := toks[3]
		other := toks[len(toks)-1]
		other = strings.TrimSuffix(other, ".")
		val, err := strconv.Atoi(valueStr)
		if err != nil {
			continue
		}
		if gainOrLose == "lose" {
			val = -val
		}
		if hap[name] == nil {
			hap[name] = make(map[string]int)
		}
		hap[name][other] = val
		namesMap[name] = struct{}{}
		namesMap[other] = struct{}{}
	}

	// build names slice
	names := make([]string, 0, len(namesMap))
	for n := range namesMap {
		names = append(names, n)
	}

	if len(names) == 0 {
		return "0"
	}

	// permutation search for best total happiness (circular)
	best := int64(-1 << 60)
	var permute func(int)
	permute = func(i int) {
		if i == len(names)-1 {
			// compute total
			var total int64
			n := len(names)
			for k := 0; k < n; k++ {
				me := names[k]
				left := names[(k-1+n)%n]
				right := names[(k+1)%n]
				total += int64(hap[me][left])
				total += int64(hap[me][right])
			}
			if total > best {
				best = total
			}
			return
		}
		for j := i; j < len(names); j++ {
			names[i], names[j] = names[j], names[i]
			permute(i + 1)
			names[i], names[j] = names[j], names[i]
		}
	}
	permute(0)

	return strconv.FormatInt(best, 10)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	hap := make(map[string]map[string]int)
	namesMap := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		toks := strings.Fields(line)
		if len(toks) < 11 {
			continue
		}
		name := toks[0]
		gainOrLose := toks[2]
		valueStr := toks[3]
		other := toks[len(toks)-1]
		other = strings.TrimSuffix(other, ".")
		val, err := strconv.Atoi(valueStr)
		if err != nil {
			continue
		}
		if gainOrLose == "lose" {
			val = -val
		}
		if hap[name] == nil {
			hap[name] = make(map[string]int)
		}
		hap[name][other] = val
		namesMap[name] = struct{}{}
		namesMap[other] = struct{}{}
	}

	// add "You" with zero relationships to everyone and vice versa
	const you = "You"
	if hap[you] == nil {
		hap[you] = make(map[string]int)
	}
	for n := range namesMap {
		hap[you][n] = 0
		if hap[n] == nil {
			hap[n] = make(map[string]int)
		}
		hap[n][you] = 0
	}
	namesMap[you] = struct{}{}

	// build names slice
	names := make([]string, 0, len(namesMap))
	for n := range namesMap {
		names = append(names, n)
	}

	if len(names) == 0 {
		return "0"
	}

	// permutation search for best total happiness (circular)
	best := int64(-1 << 60)
	var permute func(int)
	permute = func(i int) {
		if i == len(names)-1 {
			// compute total
			var total int64
			n := len(names)
			for k := 0; k < n; k++ {
				me := names[k]
				left := names[(k-1+n)%n]
				right := names[(k+1)%n]
				total += int64(hap[me][left])
				total += int64(hap[me][right])
			}
			if total > best {
				best = total
			}
			return
		}
		for j := i; j < len(names); j++ {
			names[i], names[j] = names[j], names[i]
			permute(i + 1)
			names[i], names[j] = names[j], names[i]
		}
	}
	permute(0)

	return strconv.FormatInt(best, 10)
}

func New() *Puzzle {
	return &Puzzle{}
}
