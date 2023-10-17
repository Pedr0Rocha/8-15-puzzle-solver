package puzzle

import "github.com/Pedr0Rocha/8-15-puzzle-solver/board"

var Puzzle1 = Puzzle{
	InitialBoard: board.Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	},
	MinSteps: 1,
}

var Puzzle5 = Puzzle{
	InitialBoard: board.Board{
		{1, 2, 3},
		{5, 6, 0},
		{4, 7, 8},
	},
	MinSteps: 5,
}

var Puzzle7 = Puzzle{
	InitialBoard: board.Board{
		{4, 1, 3},
		{7, 2, 5},
		{8, 0, 6},
	},
	MinSteps: 7,
}

var Puzzle13 = Puzzle{
	InitialBoard: board.Board{
		{4, 3, 1},
		{0, 7, 2},
		{8, 5, 6},
	},
	MinSteps: 13,
}

var Puzzle20 = Puzzle{
	InitialBoard: board.Board{
		{7, 4, 3},
		{2, 8, 6},
		{0, 5, 1},
	},
	MinSteps: 20,
}

var Puzzle26 = Puzzle{
	InitialBoard: board.Board{
		{4, 8, 7},
		{5, 3, 1},
		{0, 6, 2},
	},
	MinSteps: 26,
}

var Puzzle31 = Puzzle{
	InitialBoard: board.Board{
		{6, 4, 7},
		{8, 5, 0},
		{3, 2, 1},
	},
	MinSteps: 31,
}
