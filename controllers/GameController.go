package controllers

import (
	"net/http"

	"github.com/ahmad2smile/battle-gin/context"
	"github.com/ahmad2smile/battle-gin/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GamesController(router *gin.Engine) {

	group := router.Group("/games")
	{
		group.GET("/", getAll)
		group.POST("/", newGame)
	}
}

func getAll(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func newGame(ctx *gin.Context) {
	var gameDto models.GameDTO

	ctx.ShouldBindJSON(&gameDto)

	game := models.Game{ID: uuid.NewV4(), PlayerRole: gameDto.PlayerRole}

	if err := context.Db.Create(game).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, game)
	}
}
