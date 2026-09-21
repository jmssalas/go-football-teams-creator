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
	var season SeasonQuery

	if err := c.ShouldBindQuery(&season); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	seasonID := season.SeasonID
	if seasonID == 0 {
		currentSeason, err := h.dbAdapter.GetLastSeason()
		if err != nil {
			Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
			return
		}
		seasonID = currentSeason.ID
	}

	dbPlayers, err := h.dbAdapter.GetPlayers(seasonID)

	if err != nil {
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	players := make([]Player, len(dbPlayers))
	for i, row := range dbPlayers {
		players[i] = Player{
			ID:                row.ID,
			Name:              row.Name,
			MatchesWon:        row.MatchesWon,
			MatchesDrawn:      row.MatchesDrawn,
			MatchesLost:       row.MatchesLost,
			GoalsFor:          row.GoalsFor,
			GoalsAgainst:      row.GoalsAgainst,
			TotalMatches:      row.TotalMatches,
			TotalPoints:       row.TotalPoints,
			VictoryPercentage: row.VictoryPercentage,
		}
	}

	OK(c, players)
}

func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	var player DeletePlayerReq

	if err := c.ShouldBindUri(&player); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	if err := h.dbAdapter.DeletePlayer(player.ID); err != nil {
		// @TODO: Error handling
	}

	OK(c, nil)
}
