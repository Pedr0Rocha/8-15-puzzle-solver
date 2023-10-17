package puzzle

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/Pedr0Rocha/8-15-puzzle-solver/board"
)

func (p Puzzle) SolveAStarWithPriorityQueue() int {
	openSet := make(map[string]board.Board)
	hash := p.InitialBoard.GenerateHash()
	openSet[hash] = p.InitialBoard

	gScore := make(map[string]int)   // cost from start board to current board
	fScore := make(PriorityQueue, 1) // cost from start to goal
	gScore[hash] = 0
	fScore[0] = &PriorityQueueItem{
		boardHash: hash,
		priority:  heuristic(p.InitialBoard),
		index:     0,
	}
	heap.Init(&fScore)

	path := make(map[string]board.Board)

	for len(openSet) != 0 {
		lowestFScoreHash := heap.Pop(&fScore).(*PriorityQueueItem).boardHash
		current := openSet[lowestFScoreHash]
		if current == nil {
			continue
		}
		currentHash := lowestFScoreHash

		if IsSolution(currentHash, SolutionHash) {
			completePath := reconstructPath(path, current)

			for _, p := range completePath {
				p.Print()
			}

			fmt.Printf("Found solution in %v movements, optimal is %v\n", gScore[currentHash], p.MinSteps)
			return gScore[currentHash]
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
				newFScore := gScore[neighborHash] + heuristic(neighbor)
				newFScoreEntry := &PriorityQueueItem{
					boardHash: neighborHash,
					priority:  newFScore,
				}
				heap.Push(&fScore, newFScoreEntry)
				fScore.update(newFScoreEntry, neighborHash, newFScore)

				if _, found := openSet[neighborHash]; !found {
					openSet[neighborHash] = neighbor
				}
			}
		}
	}
	fmt.Println("Failed to find the solution")
	return -1
}
