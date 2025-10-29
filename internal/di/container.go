package di

import (
	"fmt"
	"sync"

	"github.com/teilin/aoc-go/pkg/puzzle"
)

type Container struct {
	mu       sync.RWMutex
	registry map[string]puzzle.Puzzle
}

func NewContainer() *Container {
	return &Container{
		registry: make(map[string]puzzle.Puzzle),
	}
}

func (c *Container) Register(year string, day string, puzzle puzzle.Puzzle) {
	c.mu.Lock()
	defer c.mu.Unlock()
	key := fmt.Sprintf("%s-%s", year, day)
	c.registry[key] = puzzle
}

func (c *Container) Get(year string, day string) (puzzle.Puzzle, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	key := fmt.Sprintf("%s-%s", year, day)
	puzzle, exists := c.registry[key]
	if !exists {
		return nil, fmt.Errorf("puzzle not found for year %s and day %s", year, day)
	}
	return puzzle, nil
}
