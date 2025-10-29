# Advent of Code Go Console Application

This project is a console application written in Go that allows users to solve Advent of Code puzzles from the years 2015 to 2024. The application is designed to be modular and extensible, utilizing interfaces and dependency injection to manage puzzle implementations.

## Project Structure

```
aoc-go
├── cmd
│   └── aoc
│       └── main.go          # Entry point of the application
├── internal
│   ├── app
│   │   ├── app.go           # Main application logic
│   │   └── runner.go        # Logic for running puzzles
│   ├── di
│   │   └── container.go      # Dependency injection container
│   └── puzzles
│       ├── 2015
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2015 Day 1 puzzle
│       ├── 2016
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2016 Day 1 puzzle
│       ├── 2017
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2017 Day 1 puzzle
│       ├── 2018
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2018 Day 1 puzzle
│       ├── 2019
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2019 Day 1 puzzle
│       ├── 2020
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2020 Day 1 puzzle
│       ├── 2021
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2021 Day 1 puzzle
│       ├── 2022
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2022 Day 1 puzzle
│       ├── 2023
│       │   └── day01
│       │       └── puzzle.go # Implementation of 2023 Day 1 puzzle
│       └── 2024
│           └── day01
│               └── puzzle.go # Implementation of 2024 Day 1 puzzle
├── pkg
│   └── puzzle
│       ├── puzzle.go        # Puzzle interface definition
│       └── registry.go      # Registration and lookup for puzzle implementations
├── configs
│   └── config.yaml          # Configuration settings
├── Makefile                 # Build instructions
├── go.mod                   # Module and dependencies
├── .gitignore               # Files to ignore in version control
└── README.md                # Project documentation
```

## Getting Started

To get started with the application, clone the repository and navigate to the project directory. You can build and run the application using the provided Makefile.

### Prerequisites

- Go 1.16 or later
- Git

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd aoc-go
   ```

2. Build the application:
   ```
   make build
   ```

3. Run the application:
   ```
   ./cmd/aoc/aoc <year> <day>
   ```

Replace `<year>` and `<day>` with the desired puzzle year and day.

## Usage

The application will prompt you for the year and day of the puzzle you wish to solve. It will then execute the corresponding puzzle implementation and display the results for both parts.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.