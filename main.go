package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/Pedr0Rocha/8-15-puzzle-solver/puzzle"
)

func run(puzzle puzzle.Puzzle, withProfiler bool) int {
	if withProfiler {
		f, err := os.Create("solver.prof")
		if err != nil {
			fmt.Println(err)
			return -1
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	minSteps := puzzle.SolveAStarWithPriorityQueue()
	return minSteps
}

func main() {
	run(puzzle.Puzzle31, false)
}
