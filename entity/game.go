package entity

type PlayerMove struct {
	XIndex int `json:"x_index" binding:"min=0,max=2"`
	YIndex int `json:"y_index" binding:"min=0,max=2"`
}

type Game struct {
	GameBoard Board
	GameScore Score
	Player1   Player
	Player2   Player
	Turn      int
	Moves     int
	State     int
}
type Board [3][3]string

type Score struct {
	HorizontalScore    [3]int
	VerticalScore      [3]int
	LeftDiagonalScore  int
	RightDiagonalScore int
}

type Player struct {
	ID        int
	Symbol    string
	Increment int
}

func NewGame() Game {
	game := Game{}
	game.Moves = 0
	game.GameBoard = newBoard()
	game.GameScore = newScore()
	game.Player1 = newPlayer(0, "O", 1)
	game.Player2 = newPlayer(1, "X", -1)
	game.Turn = 0
	game.State = 0
	return game
}

func newScore() Score {
	score := Score{}
	score.HorizontalScore = [3]int{0, 0, 0}
	score.VerticalScore = [3]int{0, 0, 0}
	score.LeftDiagonalScore = 0
	score.RightDiagonalScore = 0
	return score
}

func newPlayer(id int, symbol string, increment int) Player {
	player := Player{id, symbol, increment}
	return player
}

func newBoard() Board {
	board := Board(
		[3][3]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		})
	return board
}
