package db

import (
	"time"

	"gorm.io/gorm"
)

const (
	TeamA uint = iota
	TeamB
)

type Player struct {
	gorm.Model
	Name    string
	Matches []Match `gorm:"many2many:player_matches;"`
}

type PlayerStats struct {
	ID                uint    `gorm:"column:id"`
	Name              string  `gorm:"column:name"`
	MatchesWon        uint    `gorm:"column:matches_won"`
	MatchesDrawn      uint    `gorm:"column:matches_drawn"`
	MatchesLost       uint    `gorm:"column:matches_lost"`
	GoalsFor          uint    `gorm:"column:goals_for"`
	GoalsAgainst      uint    `gorm:"column:goals_against"`
	TotalMatches      uint    `gorm:"-"`
	TotalPoints       uint    `gorm:"-"`
	VictoryPercentage float64 `gorm:"-"`
}

type Season struct {
	gorm.Model
	Name string
}

type Match struct {
	gorm.Model
	Date       time.Time
	TeamAScore uint
	TeamBScore uint
	SeasonID   uint
	Season     Season
	Players    []Player `gorm:"many2many:player_matches;"`
}

type PlayerMatch struct {
	gorm.Model
	PlayerID uint `gorm:"primaryKey"`
	MatchID  uint `gorm:"primaryKey"`
	Team     uint // TeamA, TeamB
	Score    uint
}
