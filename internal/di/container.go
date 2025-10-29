package di

import (
	"fmt"
	"sync"

	"github.com/teilin/aoc-go/internal/models"
)

type Container struct {
	mu       sync.RWMutex
	registry map[string]models.Puzzle
}

func NewContainer() *Container {
	return &Container{
		registry: make(map[string]models.Puzzle),
	}
}

func (c *Container) Register(year int, day int, puzzle models.Puzzle) {
	c.mu.Lock()
	defer c.mu.Unlock()
	key := fmt.Sprintf("%d-%d", year, day)
	c.registry[key] = puzzle
}

func (c *Container) Get(year int, day int) (models.Puzzle, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	key := fmt.Sprintf("%d-%d", year, day)
	puzzle, exists := c.registry[key]
	if !exists {
		return nil, fmt.Errorf("puzzle not found for year %d and day %d", year, day)
	}
	return puzzle, nil
}
