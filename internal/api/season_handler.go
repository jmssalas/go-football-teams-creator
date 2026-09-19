package api

import (
	"football-teams-creator/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeasonHandler struct {
	dbAdapter *db.DbAdapter
}

func (h *SeasonHandler) CreateSeason(c *gin.Context) {
	var season CreateSeasonReq

	if err := c.ShouldBindJSON(&season); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	err := h.dbAdapter.CreateSeason(&db.Season{Name: season.Name})

	if err != nil {
		// @TODO: Error handling
	}

	OK(c, nil)
}

func (h *SeasonHandler) GetSeasons(c *gin.Context) {
	dbSeasons, err := h.dbAdapter.GetSeasons()

	if err != nil {
		//@TODO: Error handling
	}

	seasons := make([]Season, len(dbSeasons))
	for i, db := range dbSeasons {
		seasons[i] = Season{
			Name: db.Name,
		}
	}

	OK(c, seasons)
}

func (h *SeasonHandler) GetCurrentSeason(c *gin.Context) {
	dbSeason, err := h.dbAdapter.GetLastSeason()

	if err != nil {
		//@TODO: Error handling
	}

	OK(c, dbSeason)
}
