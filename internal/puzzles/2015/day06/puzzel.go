package day06

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day06/input.txt")
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

type instr struct {
	action string // "on", "off", "toggle"
	x1, y1 int
	x2, y2 int
}

func parseCoord(s string) (int, int, error) {
	parts := strings.Split(strings.TrimSpace(s), ",")
	if len(parts) != 2 {
		return 0, 0, strconv.ErrSyntax
	}
	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func parseInstructions(content []byte) []instr {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var ops []instr
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		var action, fromStr, toStr string

		if fields[0] == "turn" && len(fields) >= 5 {
			action = fields[1] // "on" or "off"
			fromStr = fields[2]
			toStr = fields[4]
		} else if fields[0] == "toggle" && len(fields) >= 4 {
			action = "toggle"
			fromStr = fields[1]
			toStr = fields[3]
		} else {
			continue
		}

		x1, y1, err1 := parseCoord(fromStr)
		x2, y2, err2 := parseCoord(toStr)
		if err1 != nil || err2 != nil {
			continue
		}
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		ops = append(ops, instr{action: action, x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return ops
}

const size = 1000

func applyPart1(ops []instr) int {
	grid := make([]bool, size*size)
	for _, op := range ops {
		// clamp to grid bounds just in case
		if op.x1 < 0 {
			op.x1 = 0
		}
		if op.y1 < 0 {
			op.y1 = 0
		}
		if op.x2 >= size {
			op.x2 = size - 1
		}
		if op.y2 >= size {
			op.y2 = size - 1
		}
		for y := op.y1; y <= op.y2; y++ {
			rowBase := y * size
			for x := op.x1; x <= op.x2; x++ {
				idx := rowBase + x
				switch op.action {
				case "on":
					grid[idx] = true
				case "off":
					grid[idx] = false
				case "toggle":
					grid[idx] = !grid[idx]
				}
			}
		}
	}
	count := 0
	for _, v := range grid {
		if v {
			count++
		}
	}
	return count
}

func applyPart2(ops []instr) int64 {
	grid := make([]int, size*size)
	for _, op := range ops {
		// clamp to grid bounds just in case
		if op.x1 < 0 {
			op.x1 = 0
		}
		if op.y1 < 0 {
			op.y1 = 0
		}
		if op.x2 >= size {
			op.x2 = size - 1
		}
		if op.y2 >= size {
			op.y2 = size - 1
		}
		for y := op.y1; y <= op.y2; y++ {
			rowBase := y * size
			for x := op.x1; x <= op.x2; x++ {
				idx := rowBase + x
				switch op.action {
				case "on":
					grid[idx] += 1
				case "off":
					if grid[idx] > 0 {
						grid[idx] -= 1
					}
				case "toggle":
					grid[idx] += 2
				}
			}
		}
	}
	var total int64
	for _, v := range grid {
		total += int64(v)
	}
	return total
}

func (p *Puzzle) Part1() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	ops := parseInstructions(content)
	count := applyPart1(ops)
	return strconv.Itoa(count)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	ops := parseInstructions(content)
	total := applyPart2(ops)
	return strconv.FormatInt(total, 10)
}

func New() *Puzzle {
	return &Puzzle{}
}
