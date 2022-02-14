package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tic_tac_toe/controller"
)

var (
	gameController = controller.NewGameController()
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//r.Static("/static", "./public")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/game", func(ctx *gin.Context) {
		board, state, turn := gameController.GetGame()
		ctx.JSON(200, gin.H{
			"board": board,
			"state": state,
			"turn":  turn,
		})
	})
	router.GET("/reset", func(ctx *gin.Context) {
		gameController.ResetGame()
		ctx.Status(200)
	})
	router.POST("/play", func(ctx *gin.Context) {
		result, err := gameController.Play(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, result)
	})
	fmt.Println("Server is running on port 3000")
	router.Run(":3000")
}
