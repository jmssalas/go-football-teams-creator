package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TeamsHandler struct {
	filepath string
}

func (h *TeamsHandler) CreateTeams(c *gin.Context) {
	var teams TeamsReq

	if err := c.ShouldBindJSON(&teams); err != nil {
		Fail(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	jsonBytes, err := json.Marshal(teams)
	if err != nil {
		// @TODO: Error handling
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	if err = os.WriteFile(h.filepath, jsonBytes, 0644); err != nil {
		// @TODO: Error handling
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	OK(c, nil)
}

func (h *TeamsHandler) GetTeams(c *gin.Context) {
	var teams TeamsRes

	data, err := os.ReadFile(h.filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Ignore
			OK(c, teams)
			return
		}
		// @TODO: Error handling
	}

	if err = json.Unmarshal(data, &teams); err != nil {
		// @TODO: Error handling
		Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	OK(c, teams)
}
