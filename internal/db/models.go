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
