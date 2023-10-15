package main

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

func main() {
	board := Board{
		{0, 2, 8},
		{5, 6, 1},
		{3, 7, 4},
	}
	board.Print()
	mutations := board.getMutations()

	for _, mutation := range mutations {
		mutation.Print()
	}
}
