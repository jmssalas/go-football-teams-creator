package api

import "time"

type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
	Meta    *Meta      `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type CreatePlayerReq struct {
	Name string `json:"name" binding:"required"`
}

type CreateSeasonReq struct {
	Name string `json:"name" binding:"required"`
}

type TeamReq struct {
	PlayerID uint `json:"playerId" binding:"required"`
}

type CreateMatchReq struct {
	Date       time.Time `json:"date" binding:"required"`
	SeasonID   uint      `json:"seasonId" binding:"required"`
	TeamA      []TeamReq `json:"teamA" binding:"required"`
	TeamB      []TeamReq `json:"teamB" binding:"required"`
	TeamAGoals uint      `json:"teamAGoals" binding:"required"`
	TeamBGoals uint      `json:"teamBGoals" binding:"required"`
}

type Player struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	MatchesWon        uint    `json:"matchesWon"`
	MatchesDrawn      uint    `json:"matchesDrawn"`
	MatchesLost       uint    `json:"matchesLost"`
	GoalsFor          uint    `json:"goalsFor"`
	GoalsAgainst      uint    `json:"goalsAgainst"`
	TotalMatches      uint    `json:"totalMatches"`
	TotalPoints       uint    `json:"totalPoints"`
	VictoryPercentage float64 `json:"victoryPercentage"`
}

type Season struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DeletePlayerReq struct {
	ID uint `uri:"id" binding:"required"`
}

type TeamPlayer struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	VictoryPercentage float64 `json:"victoryPercentage"`
}

type Teams struct {
	TeamA []TeamPlayer `json:"teamA"`
	TeamB []TeamPlayer `json:"teamB"`
	Date  string       `json:"date,omitempty"`
}

type TeamsReq struct {
	Teams []Teams `json:"teams"`
}

type TeamsRes TeamsReq
