package puzzle

import (
	"github.com/Pedr0Rocha/8-15-puzzle-solver/board"
)

type Puzzle struct {
	InitialBoard board.Board
	MinSteps     int
}

func IsSolution(hash string, solutionHash string) bool {
	return hash == solutionHash
}

var Solution = board.Board{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 0},
}
var SolutionHash = Solution.GenerateHash()
