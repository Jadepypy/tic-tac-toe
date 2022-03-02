package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tic_tac_toe/server/controller"
	"tic_tac_toe/server/entity"
)

var (
	gameController = controller.NewGameController()
)

func main() {
	go MainGRPC()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("./server/static/*")

	//r.Static("/static", "./static")
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
		var playerMove entity.PlayerMove
		err := ctx.ShouldBindJSON(&playerMove)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		result, err1 := gameController.Play(playerMove.XIndex, playerMove.YIndex)
		if err1 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, result)
	})
	fmt.Println("Server is running on port 3000")
	router.Run(":3000")
}
