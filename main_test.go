package main

import (
	"testing"

	"github.com/Pedr0Rocha/8-15-puzzle-solver/board"
	"github.com/Pedr0Rocha/8-15-puzzle-solver/puzzle"
)

func TestRun(t *testing.T) {
	inputs := []struct {
		name   string
		puzzle puzzle.Puzzle
	}{
		{
			name: "3x3 - 13 steps",
			puzzle: puzzle.Puzzle{
				InitialBoard: board.Board{
					{4, 3, 1},
					{0, 7, 2},
					{8, 5, 6},
				},
				MinSteps: 13,
			},
		},
		{
			name: "3x3 - 20 steps",
			puzzle: puzzle.Puzzle{
				InitialBoard: board.Board{
					{7, 4, 3},
					{2, 8, 6},
					{0, 5, 1},
				},
				MinSteps: 20,
			},
		},
	}

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			steps := run(tt.puzzle, false)

			if steps != tt.puzzle.MinSteps {
				t.Errorf("Wrong min. steps. got:%v, expected:%v", steps, tt.puzzle.MinSteps)
			}
		})
	}
}
