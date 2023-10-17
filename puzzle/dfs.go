package puzzle

import (
	"fmt"

	"github.com/Pedr0Rocha/8-15-puzzle-solver/board"
)

const (
	MAX_STEPS_3X3 = 181_440
)

func (p Puzzle) SolveDFS() int {
	fmt.Println("Trying to solve puzzle:")
	p.InitialBoard.Print()

	tries := 0

	var queue []board.Board
	visited := make(map[string]bool)
	queue = append(queue, p.InitialBoard)
	for len(queue) != 0 {
		curBoard := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		boardHash := curBoard.GenerateHash()

		_, found := visited[boardHash]
		if found {
			continue
		}

		visited[boardHash] = true

		if IsSolution(boardHash, SolutionHash) {
			fmt.Printf("It took %v tries\n", tries)
			return tries
		}

		queue = append(queue, curBoard.GetNeighbors()...)
		tries += 1

		if tries >= MAX_STEPS_3X3 {
			fmt.Println("Exausted attempts to solve puzzle")
			break
		}
	}
	fmt.Println("Didn't solve it")
	return -1
}
