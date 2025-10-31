package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/teilin/aoc-go/internal/di"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day01"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day02"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day03"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day04"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day05"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day06"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day07"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day08"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day09"
	"github.com/teilin/aoc-go/internal/puzzles/2015/day10"

	"github.com/teilin/aoc-go/internal/app"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: aoc <year> <day>")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil || year < 2015 || year > 2024 {
		fmt.Println("Invalid year. Please enter a year between 2015 and 2024.")
		return
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Invalid day. Please enter a day between 1 and 25.")
		return
	}

	container := di.NewContainer()
	app := app.App{
		PuzzleRegistry: *container,
	}
	app.PuzzleRegistry.Register(2015, 1, day01.New())
	app.PuzzleRegistry.Register(2015, 2, day02.New())
	app.PuzzleRegistry.Register(2015, 3, day03.New())
	app.PuzzleRegistry.Register(2015, 4, day04.New())
	app.PuzzleRegistry.Register(2015, 5, day05.New())
	app.PuzzleRegistry.Register(2015, 6, day06.New())
	app.PuzzleRegistry.Register(2015, 7, day07.New())
	app.PuzzleRegistry.Register(2015, 8, day08.New())
	app.PuzzleRegistry.Register(2015, 9, day09.New())
	app.PuzzleRegistry.Register(2015, 10, day10.New())

	app.Run(year, day)
}
