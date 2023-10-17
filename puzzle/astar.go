package puzzle

import (
	"fmt"
	"math"

	"github.com/Pedr0Rocha/8-15-puzzle-solver/board"
)

func (p Puzzle) SolveAStar() int {
	openSet := make(map[string]board.Board)
	hash := p.InitialBoard.GenerateHash()
	openSet[hash] = p.InitialBoard

	gScore := make(map[string]int) // cost from start board to current board
	fScore := make(map[string]int) // cost from start to goal
	gScore[hash] = 0
	fScore[hash] = heuristic(p.InitialBoard)

	path := make(map[string]board.Board)

	for len(openSet) != 0 {
		lowestFScoreHash := getLowestFScoreHash(openSet, fScore)
		current := openSet[lowestFScoreHash]
		currentHash := lowestFScoreHash

		if IsSolution(currentHash, SolutionHash) {
			completePath := reconstructPath(path, current)

			for _, p := range completePath {
				p.Print()
			}

			fmt.Printf("Found solution in %v movements, optimal is %v\n", gScore[currentHash], p.MinSteps)
			return fScore[currentHash]
		}

		delete(openSet, currentHash)

		for _, neighbor := range current.GetNeighbors() {
			neighborHash := neighbor.GenerateHash()
			tentativeGScore := gScore[currentHash] + 1 // + 1 because its the distance between current and neighbor

			if _, found := gScore[neighborHash]; !found {
				gScore[neighborHash] = math.MaxInt
			}

			if tentativeGScore < gScore[neighborHash] {
				path[neighborHash] = current
				gScore[neighborHash] = tentativeGScore
				fScore[neighborHash] = gScore[neighborHash] + heuristic(neighbor)

				if _, found := openSet[neighborHash]; !found {
					openSet[neighborHash] = neighbor
				}
			}
		}
	}
	fmt.Println("Failed to find the solution")
	return -1
}

func getLowestFScoreHash(openSet map[string]board.Board, fScore map[string]int) string {
	lowest := math.MaxInt
	lowestHash := ""
	for hash, score := range fScore {
		if score < lowest {
			if _, found := openSet[hash]; !found {
				continue
			}
			lowest = score
			lowestHash = hash
		}
	}
	return lowestHash
}
