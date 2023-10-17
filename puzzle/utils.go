package puzzle

import "github.com/Pedr0Rocha/8-15-puzzle-solver/board"

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Manhattan distance
func heuristic(b board.Board) int {
	distance := 0
	for row := range b {
		for col := range b[row] {
			if b[row][col] != 0 {
				targetRow := (b[row][col] - 1) / 3
				targetCol := (b[row][col] - 1) % 3
				distance += abs(row-targetRow) + abs(col-targetCol)
			}
		}
	}
	return distance
}

func isInArray(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func reconstructPath(path map[string]board.Board, current board.Board) []board.Board {
	completePath := []board.Board{current}

	keys := make([]string, 0, len(path))
	for k := range path {
		keys = append(keys, k)
	}

	for {
		hash := current.GenerateHash()
		if !isInArray(keys, hash) {
			break
		}

		current = path[hash]
		completePath = append(completePath, current)
	}

	// reverse the path to display in correct order
	for i, j := 0, len(completePath)-1; i < j; i, j = i+1, j-1 {
		completePath[i], completePath[j] = completePath[j], completePath[i]
	}

	return completePath
}
