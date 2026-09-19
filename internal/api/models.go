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
	PlayerID uint `json:"playerId"`
	Score    uint `json:"score"`
}

type CreateMatchReq struct {
	Date     time.Time `json:"date"`
	SeasonID uint      `json:"seasonId" binding:"required"`
	TeamA    []TeamReq `json:"teamA" binding:"required"`
	TeamB    []TeamReq `json:"teamB" binding:"required"`
}

type Player struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Season struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
