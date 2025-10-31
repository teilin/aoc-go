package day09

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct{}

func readInput() ([]byte, error) {
	file, err := os.Open("./internal/puzzles/2015/day09/input.txt")
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
	dist, nodes := parseDistances(content)
	min, ok := bestRoute(dist, nodes, false)
	if !ok {
		return ""
	}
	return strconv.Itoa(min)
}

func (p *Puzzle) Part2() string {
	content, err := readInput()
	if err != nil {
		return ""
	}
	dist, nodes := parseDistances(content)
	max, ok := bestRoute(dist, nodes, true)
	if !ok {
		return ""
	}
	return strconv.Itoa(max)
}

func New() *Puzzle {
	return &Puzzle{}
}

func parseDistances(content []byte) (map[string]map[string]int, []string) {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	dist := make(map[string]map[string]int)
	nodesMap := make(map[string]struct{})
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue
		}
		d, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		ends := strings.Split(parts[0], " to ")
		if len(ends) != 2 {
			continue
		}
		a := strings.TrimSpace(ends[0])
		b := strings.TrimSpace(ends[1])
		if dist[a] == nil {
			dist[a] = make(map[string]int)
		}
		if dist[b] == nil {
			dist[b] = make(map[string]int)
		}
		dist[a][b] = d
		dist[b][a] = d
		nodesMap[a] = struct{}{}
		nodesMap[b] = struct{}{}
	}

	nodes := make([]string, 0, len(nodesMap))
	for n := range nodesMap {
		nodes = append(nodes, n)
	}
	return dist, nodes
}

// routeDistance returns the distance for the current permutation of nodes,
// and a boolean indicating whether the route is valid (all consecutive pairs have a distance).
func routeDistance(nodes []string, dist map[string]map[string]int) (int, bool) {
	sum := 0
	for k := 0; k+1 < len(nodes); k++ {
		a := nodes[k]
		b := nodes[k+1]
		dmap, ok := dist[a]
		if !ok {
			return 0, false
		}
		d, ok2 := dmap[b]
		if !ok2 {
			return 0, false
		}
		sum += d
	}
	return sum, true
}

// bestRoute computes either the minimum (findMax=false) or maximum (findMax=true) route distance.
func bestRoute(dist map[string]map[string]int, nodes []string, findMax bool) (int, bool) {
	if len(nodes) == 0 {
		return 0, false
	}
	best := 0
	first := true
	var permute func(int)
	permute = func(i int) {
		if i == len(nodes)-1 {
			if d, ok := routeDistance(nodes, dist); ok {
				if first {
					best = d
					first = false
				} else {
					if findMax {
						if d > best {
							best = d
						}
					} else {
						if d < best {
							best = d
						}
					}
				}
			}
			return
		}
		for j := i; j < len(nodes); j++ {
			nodes[i], nodes[j] = nodes[j], nodes[i]
			permute(i + 1)
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	permute(0)
	if first {
		return 0, false
	}
	return best, true
}
