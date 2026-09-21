package api

import (
	"net/http"

	"football-teams-creator/db"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}

func Router(dbAdapter *db.DbAdapter, teamsFilePath string) *gin.Engine {
	playerHandler := PlayerHandler{dbAdapter: dbAdapter}
	seasonHandler := SeasonHandler{dbAdapter: dbAdapter}
	matchHandler := MatchHandler{dbAdapter: dbAdapter}
	teamsHandler := TeamsHandler{filepath: teamsFilePath}

	r := gin.Default()

	r.POST("/api/players", playerHandler.CreatePlayer)
	r.GET("/api/players", playerHandler.GetPlayers)
	r.DELETE("/api/players/:id", playerHandler.DeletePlayer)

	r.POST("/api/seasons", seasonHandler.CreateSeason)
	r.GET("/api/seasons", seasonHandler.GetSeasons)
	r.GET("/api/seasons/current", seasonHandler.GetCurrentSeason)

	r.POST("/api/matches", matchHandler.CreateMatch)

	r.GET("/api/teams", teamsHandler.GetTeams)
	r.POST("/api/teams", teamsHandler.CreateTeams)

	return r
}
