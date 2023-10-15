package main

import (
	"fmt"
)

type Puzzle struct {
	InitialBoard    Board
	Solution        Board
	StepsToSolution int
}

type Board [][]int

type Position struct {
	row int
	col int
}

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

var puzzleEasy = Puzzle{
	InitialBoard: Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	},
	Solution: Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	},
	StepsToSolution: 1,
}

func IsSolution(test Board, solution Board) bool {
	for row := range solution {
		for col := range solution[row] {
			if solution[row][col] != test[row][col] {
				return false
			}
		}
	}
	return true
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

func (p Puzzle) Solve() bool {
	fmt.Println("Trying to solve puzzle:")
	p.InitialBoard.Print()

	maxTries := 181_440

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

		if IsSolution(curBoard, p.Solution) {
			fmt.Printf("It took %v tries\n", 181_440-maxTries)
			return true
		}

		queue = append(queue, curBoard.getMutations()...)
		maxTries = maxTries - 1

		if maxTries <= 0 {
			break
		}
	}
	fmt.Println("Didn't solve it")
	return false
}

func main() {
	puzzleEasy.Solve()
}
