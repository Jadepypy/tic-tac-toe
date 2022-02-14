package controller

import (
	"github.com/gin-gonic/gin"
	"tic_tac_toe/entity"
)

//Game State
const (
	Ongoing = iota
	Draw
	Win
	Invalid
)

type GameController interface {
	Play(ctx *gin.Context) (map[string]int, error)
	ResetGame()
	GetGame() ([3][3]string, int, int)
}

type gameController struct {
	game entity.Game
}

func NewGameController() GameController {
	return &gameController{entity.NewGame()}
}

func (controller *gameController) Play(ctx *gin.Context) (map[string]int, error) {
	var playerMove entity.PlayerMove
	err := ctx.ShouldBindJSON(&playerMove)
	if err != nil {
		return nil, err
	}
	x, y := playerMove.XIndex, playerMove.YIndex
	game := &controller.game
	board := &game.GameBoard
	result := make(map[string]int, 2)
	result["player"] = game.Turn

	//check if move is invalid
	if x > 2 || y > 2 || x < 0 || y < 0 {
		result["game_state"] = Invalid
		return result, nil
	} else if board[x][y] != "_" {
		result["game_state"] = Invalid
		return result, nil
	}

	//switch to next user turn
	player := entity.Player{}
	switch game.Turn {
	case game.Player1.ID:
		player = game.Player1
		game.Turn = game.Player2.ID
	case game.Player2.ID:
		player = game.Player2
		game.Turn = game.Player1.ID
	}

	board[x][y] = player.Symbol
	game.Moves += 1

	//check if the player wins the game
	score := &game.GameScore
	score.HorizontalScore[x] += player.Increment
	score.VerticalScore[y] += player.Increment
	if x == y {
		score.LeftDiagonalScore += player.Increment
	}
	if x+y == 2 {
		score.RightDiagonalScore += player.Increment
	}
	if score.HorizontalScore[x] == player.Increment*3 ||
		score.VerticalScore[y] == player.Increment*3 ||
		score.LeftDiagonalScore == player.Increment*3 ||
		score.RightDiagonalScore == player.Increment*3 {
		result["game_state"] = Win
		game.State = Win
		game.Turn = result["player"]
		return result, nil
	}

	//if no one wins and the board is full
	if game.Moves == 9 {
		result["game_state"] = Draw
		game.State = Draw
		return result, nil
	}
	result["game_state"] = Ongoing
	result["player"] = game.Turn
	return result, nil
}

func (controller *gameController) ResetGame() {
	controller.game = entity.NewGame()
}

func (controller *gameController) GetGame() ([3][3]string, int, int) {
	game := &controller.game
	return game.GameBoard, game.State, game.Turn
}
