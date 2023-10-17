package board

import "fmt"

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

func (b Board) GetNeighbors() []Board {
	pos := b.findZeroPosition()
	size := len(b)
	var neighbors []Board

	if pos.col+1 < size {
		neighbors = append(
			neighbors,
			b.swapPosition(pos, Position{row: pos.row, col: pos.col + 1}),
		)
	}

	if pos.col-1 >= 0 {
		neighbors = append(
			neighbors,
			b.swapPosition(pos, Position{row: pos.row, col: pos.col - 1}),
		)
	}

	if pos.row+1 < size {
		neighbors = append(
			neighbors,
			b.swapPosition(pos, Position{row: pos.row + 1, col: pos.col}),
		)
	}

	if pos.row-1 >= 0 {
		neighbors = append(
			neighbors,
			b.swapPosition(pos, Position{row: pos.row - 1, col: pos.col}),
		)
	}

	return neighbors
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
