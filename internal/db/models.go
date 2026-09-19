package db

import (
	"time"

	"gorm.io/gorm"
)

const (
	TeamA int = iota
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
	Date     time.Time
	AScore   int
	BScore   int
	SeasonID int
	Season   Season
	Players  []Player `gorm:"many2many:player_matches;"`
}

type PlayerMatch struct {
	gorm.Model
	PlayerID int `gorm:"primaryKey"`
	MatchID  int `gorm:"primaryKey"`
	Team     int // TeamA, TeamB
	Score    int
}
