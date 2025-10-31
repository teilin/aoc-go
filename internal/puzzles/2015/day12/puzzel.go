package day12

import (
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day12/input.txt")
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
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(string(content), -1)
	var sum int64
	for _, m := range matches {
		v, _ := strconv.ParseInt(m, 10, 64)
		sum += v
	}
	return strconv.FormatInt(sum, 10)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	var v interface{}
	if err := json.Unmarshal(content, &v); err != nil {
		return ""
	}
	sum := sumIgnoringRed(v)
	return strconv.FormatInt(sum, 10)
}

func New() *Puzzle {
	return &Puzzle{}
}

func sumIgnoringRed(v interface{}) int64 {
	switch vv := v.(type) {
	case float64:
		return int64(vv)
	case []interface{}:
		var s int64
		for _, e := range vv {
			s += sumIgnoringRed(e)
		}
		return s
	case map[string]interface{}:
		// if any property value is the string "red", ignore this object entirely
		for _, val := range vv {
			if str, ok := val.(string); ok && str == "red" {
				return 0
			}
		}
		var s int64
		for _, val := range vv {
			s += sumIgnoringRed(val)
		}
		return s
	default:
		return 0
	}
}
