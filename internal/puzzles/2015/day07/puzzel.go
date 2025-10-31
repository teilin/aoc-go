package day07

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day07/input.txt")
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
	ins := parseInstructions(content)
	signal := evalCircuit(ins, "a")
	return strconv.Itoa(int(signal))
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	ins := parseInstructions(content)

	signalA := evalCircuit(ins, "a")
	ins["b"] = instruction{op: "ASSIGN", a: strconv.Itoa(int(signalA))}
	signal := evalCircuit(ins, "a")
	return strconv.Itoa(int(signal))
}

func New() *Puzzle {
	return &Puzzle{}
}

type instruction struct {
	op   string // "ASSIGN", "AND", "OR", "LSHIFT", "RSHIFT", "NOT"
	a, b string // operands (one or two); may be numeric literals or wire names
}

func parseInstructions(content []byte) map[string]instruction {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	ins := make(map[string]instruction)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			continue
		}
		left := strings.TrimSpace(parts[0])
		target := strings.TrimSpace(parts[1])

		toks := strings.Fields(left)
		switch {
		case len(toks) == 1:
			// "123" or "x"
			ins[target] = instruction{op: "ASSIGN", a: toks[0]}
		case len(toks) == 2 && toks[0] == "NOT":
			// "NOT x"
			ins[target] = instruction{op: "NOT", a: toks[1]}
		case len(toks) == 3:
			// "x AND y", "p LSHIFT 2", etc.
			switch toks[1] {
			case "AND":
				ins[target] = instruction{op: "AND", a: toks[0], b: toks[2]}
			case "OR":
				ins[target] = instruction{op: "OR", a: toks[0], b: toks[2]}
			case "LSHIFT":
				ins[target] = instruction{op: "LSHIFT", a: toks[0], b: toks[2]}
			case "RSHIFT":
				ins[target] = instruction{op: "RSHIFT", a: toks[0], b: toks[2]}
			}
		}
	}
	return ins
}

// evalCircuit computes the signal for wire "a" using recursion + memoization.
func evalCircuit(ins map[string]instruction, target string) uint16 {
	cache := make(map[string]uint16)
	seen := make(map[string]bool)

	var evalOperand func(string) uint16
	var evalWire func(string) uint16

	evalOperand = func(tok string) uint16 {
		// numeric literal?
		if tok == "" {
			return 0
		}
		if tok[0] >= '0' && tok[0] <= '9' {
			v, _ := strconv.ParseUint(tok, 10, 16)
			return uint16(v)
		}
		return evalWire(tok)
	}

	evalWire = func(w string) uint16 {
		// memoized
		if v, ok := cache[w]; ok {
			return v
		}
		// detect loops (shouldn't happen in valid input)
		if seen[w] {
			return 0
		}
		seen[w] = true

		inst, ok := ins[w]
		var out uint16
		if !ok {
			// no instruction: maybe a numeric literal referenced as a wire (shouldn't happen)
			out = 0
		} else {
			switch inst.op {
			case "ASSIGN":
				out = evalOperand(inst.a)
			case "NOT":
				out = ^evalOperand(inst.a)
			case "AND":
				out = evalOperand(inst.a) & evalOperand(inst.b)
			case "OR":
				out = evalOperand(inst.a) | evalOperand(inst.b)
			case "LSHIFT":
				shift, _ := strconv.Atoi(inst.b)
				out = evalOperand(inst.a) << uint(shift)
			case "RSHIFT":
				shift, _ := strconv.Atoi(inst.b)
				out = evalOperand(inst.a) >> uint(shift)
			default:
				out = 0
			}
		}
		cache[w] = out
		return out
	}

	// if target is numeric literal, return it
	if target != "" && target[0] >= '0' && target[0] <= '9' {
		v, _ := strconv.ParseUint(target, 10, 16)
		return uint16(v)
	}
	return evalWire(target)
}
