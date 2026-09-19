package api

import (
	"football-teams-creator/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	dbAdapter *db.DbAdapter
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var match CreateMatchReq

	if err := c.ShouldBindJSON(&match); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	dbMatch := db.Match{
		Date:       match.Date,
		SeasonID:   match.SeasonID,
		TeamAScore: 0,
		TeamBScore: 0,
	}

	for _, p := range match.TeamA {
		dbMatch.TeamAScore += p.Score
	}

	for _, p := range match.TeamB {
		dbMatch.TeamBScore += p.Score
	}

	if err := h.dbAdapter.CreateMatch(&dbMatch); err != nil {
		// @TODO: Error handling
	}

	playerMatches := make([]db.PlayerMatch, len(match.TeamA)+len(match.TeamB))
	i := 0
	for _, p := range match.TeamA {
		playerMatches[i] = db.PlayerMatch{
			PlayerID: p.PlayerID,
			MatchID:  dbMatch.ID,
			Team:     db.TeamA,
			Score:    p.Score,
		}
		i++
	}

	for _, p := range match.TeamB {
		playerMatches[i] = db.PlayerMatch{
			PlayerID: p.PlayerID,
			MatchID:  dbMatch.ID,
			Team:     db.TeamB,
			Score:    p.Score,
		}
		i++
	}

	if err := h.dbAdapter.CreatePlayerMatches(&playerMatches); err != nil {
		// @TODO: Error handling
		// @TODO: What happens with the created match ?
	}

	OK(c, nil)
}
