package api

import (
	"football-teams-creator/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	dbAdapter *db.DbAdapter
}

func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var player CreatePlayerReq

	if err := c.ShouldBindJSON(&player); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	err := h.dbAdapter.CreatePlayer(&db.Player{Name: player.Name})

	if err != nil {
		// @TODO: Error handling
	}

	OK(c, nil)
}

func (h *PlayerHandler) GetPlayers(c *gin.Context) {
	dbPlayers, err := h.dbAdapter.GetPlayers()

	if err != nil {
		//@TODO: Error handling
	}

	players := make([]Player, len(dbPlayers))
	for i, dbPlayer := range dbPlayers {
		players[i] = Player{
			Name: dbPlayer.Name,
		}
	}

	OK(c, players)
}
