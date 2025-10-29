# Makefile

.PHONY: build run clean

build:
	go build -o aoc ./cmd/aoc

run: build
	./aoc

clean:
	go clean
	rm -f aoc