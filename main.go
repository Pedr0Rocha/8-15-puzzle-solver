package main

import (
	"fmt"
	"math"
)

type Puzzle struct {
	InitialBoard Board
	MinSteps     int
}

type Board [][]int

type Position struct {
	row int
	col int
}

const (
	MAX_STEPS_3X3 = 181_440
)

var Solution = Board{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 0},
}
var SolutionHash = Solution.GenerateHash()

func NewBoard(size int) Board {
	board := make([][]int, size)

	for i := range board {
		board[i] = make([]int, size)
	}

	return board
}

func (b Board) Print() {
	for _, row := range b {
		fmt.Println(row)
	}
	fmt.Println()
}

func (b Board) findZeroPosition() Position {
	for row := range b {
		for col := range b[row] {
			if b[row][col] == 0 {
				return Position{row: row, col: col}
			}
		}
	}
	return Position{}
}

func (b Board) copyBoard() Board {
	newBoard := NewBoard(len(b))
	for i := range b {
		for j := range b[i] {
			newBoard[i][j] = b[i][j]
		}
	}
	return newBoard
}

func (b Board) swapPosition(curPos Position, newPos Position) Board {
	boardCopy := b.copyBoard()
	boardCopy[curPos.row][curPos.col] = boardCopy[newPos.row][newPos.col]
	boardCopy[newPos.row][newPos.col] = 0

	return boardCopy
}

func (b Board) getMutations() []Board {
	pos := b.findZeroPosition()
	size := len(b)
	var mutations []Board

	if pos.col+1 < size {
		mutations = append(
			mutations,
			b.swapPosition(pos, Position{row: pos.row, col: pos.col + 1}),
		)
	}

	if pos.col-1 >= 0 {
		mutations = append(
			mutations,
			b.swapPosition(pos, Position{row: pos.row, col: pos.col - 1}),
		)
	}

	if pos.row+1 < size {
		mutations = append(
			mutations,
			b.swapPosition(pos, Position{row: pos.row + 1, col: pos.col}),
		)
	}

	if pos.row-1 >= 0 {
		mutations = append(
			mutations,
			b.swapPosition(pos, Position{row: pos.row - 1, col: pos.col}),
		)
	}

	return mutations
}

var puzzle1 = Puzzle{
	InitialBoard: Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	},
	MinSteps: 1,
}

var puzzle5 = Puzzle{
	InitialBoard: Board{
		{1, 2, 3},
		{5, 6, 0},
		{4, 7, 8},
	},
	MinSteps: 5,
}

var puzzle7 = Puzzle{
	InitialBoard: Board{
		{4, 1, 3},
		{7, 2, 5},
		{8, 0, 6},
	},
	MinSteps: 7,
}

var puzzle13 = Puzzle{
	InitialBoard: Board{
		{4, 3, 1},
		{0, 7, 2},
		{8, 5, 6},
	},
	MinSteps: 13,
}

var puzzle20 = Puzzle{
	InitialBoard: Board{
		{7, 4, 3},
		{2, 8, 6},
		{0, 5, 1},
	},
	MinSteps: 20,
}

var puzzle26 = Puzzle{
	InitialBoard: Board{
		{4, 8, 7},
		{5, 3, 1},
		{0, 6, 2},
	},
	MinSteps: 26,
}

var puzzle31 = Puzzle{
	InitialBoard: Board{
		{6, 4, 7},
		{8, 5, 0},
		{3, 2, 1},
	},
	MinSteps: 31,
}

func IsSolution(hash string) bool {
	return hash == SolutionHash
}

func (b Board) GenerateHash() string {
	var hash string
	for row := range b {
		for col := range b[row] {
			hash += fmt.Sprintf("%d ", b[row][col])
		}
	}
	return hash
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Manhattan distance
func heuristic(b Board) int {
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

func getLowestFScoreHash(openSet map[string]Board, fScore map[string]int) string {
	lowest := math.MaxInt
	lowestHash := ""
	for k, v := range fScore {
		if v < lowest {
			if _, found := openSet[k]; !found {
				continue
			}
			lowest = v
			lowestHash = k
		}
	}
	return lowestHash
}

func IsInArray(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func reconstructPath(path map[string]Board, current Board) []Board {
	completePath := []Board{current}

	keys := make([]string, 0, len(path))
	for k := range path {
		keys = append(keys, k)
	}

	for {
		hash := current.GenerateHash()
		if !IsInArray(keys, hash) {
			break
		}

		current = path[hash]
		completePath = append(completePath, current)
	}

	for i, j := 0, len(completePath)-1; i < j; i, j = i+1, j-1 {
		completePath[i], completePath[j] = completePath[j], completePath[i]
	}

	return completePath
}

func (p Puzzle) SolveAStar() int {
	openSet := make(map[string]Board)
	hash := p.InitialBoard.GenerateHash()
	openSet[hash] = p.InitialBoard

	gScore := make(map[string]int) // cost from start board to current board
	fScore := make(map[string]int) // cost from start to goal
	gScore[hash] = 0
	fScore[hash] = heuristic(p.InitialBoard)

	path := make(map[string]Board)

	for {
		if len(openSet) == 0 {
			break
		}

		lowestFScoreHash := getLowestFScoreHash(openSet, fScore)
		current := openSet[lowestFScoreHash]
		currentHash := lowestFScoreHash

		if IsSolution(currentHash) {
			completePath := reconstructPath(path, current)

			for _, p := range completePath {
				p.Print()
			}

			fmt.Printf("Found solution in %v movements, optimal is %v\n", gScore[currentHash], p.MinSteps)
			return fScore[currentHash]
		}

		delete(openSet, currentHash)

		for _, neighbor := range current.getMutations() {
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

func (p Puzzle) Solve() int {
	fmt.Println("Trying to solve puzzle:")
	p.InitialBoard.Print()

	tries := 0

	var queue []Board
	visited := make(map[string]bool)
	queue = append(queue, p.InitialBoard)
	for {
		if len(queue) == 0 {
			break
		}
		curBoard := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		boardHash := curBoard.GenerateHash()

		_, found := visited[boardHash]
		if found {
			continue
		}

		visited[boardHash] = true

		if IsSolution(boardHash) {
			fmt.Printf("It took %v tries\n", tries)
			return tries
		}

		queue = append(queue, curBoard.getMutations()...)
		tries += 1

		if tries >= MAX_STEPS_3X3 {
			fmt.Println("Exausted attempts to solve puzzle")
			break
		}
	}
	fmt.Println("Didn't solve it")
	return -1
}

func main() {
	// puzzleEasy.SolveAStar()
	puzzle31.SolveAStar()
	// puzzleHard.SolveAStar()
}
