package db

import (
	"context"
	"database/sql"
	"log"
	"math"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbAdapter struct {
	db *gorm.DB
}

func (a *DbAdapter) InitDb(filepath string) {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&Player{}, &Season{}, &Match{}, &PlayerMatch{})

	a.db = db
}

func (a *DbAdapter) CloseDb() {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
}

// Players
func (a *DbAdapter) CreatePlayer(player *Player) error {
	ctx := context.Background()
	return gorm.G[Player](a.db).Create(ctx, player)
}

func (a *DbAdapter) GetPlayers() ([]PlayerStats, error) {
	ctx := context.Background()
	var players []PlayerStats

	err := a.db.WithContext(ctx).Raw(`
		SELECT
			players.id,
			players.name,
			COALESCE(SUM(CASE
				WHEN (player_matches.team = @teamA AND matches.team_a_score > matches.team_b_score) OR
					(player_matches.team = @teamB AND matches.team_b_score > matches.team_a_score)
				THEN 1 ELSE 0
			END), 0) AS matches_won,
			COALESCE(SUM(CASE
				WHEN matches.team_a_score = matches.team_b_score THEN 1
				ELSE 0
			END), 0) AS matches_drawn,
			COALESCE(SUM(CASE
				WHEN (player_matches.team = @teamA AND matches.team_a_score < matches.team_b_score) OR
					(player_matches.team = @teamB AND matches.team_b_score < matches.team_a_score)
				THEN 1 ELSE 0
			END), 0) AS matches_lost,
			COALESCE(SUM(CASE
				WHEN player_matches.team = @teamA THEN matches.team_a_score
				WHEN player_matches.team = @teamB THEN matches.team_b_score
				ELSE 0
			END), 0) AS goals_for,
			COALESCE(SUM(CASE
				WHEN player_matches.team = @teamA THEN matches.team_b_score
				WHEN player_matches.team = @teamB THEN matches.team_a_score
				ELSE 0
			END), 0) AS goals_against
		FROM players
		LEFT JOIN player_matches ON players.id = player_matches.player_id
		LEFT JOIN matches ON matches.id = player_matches.match_id
		GROUP BY players.id, players.name
		ORDER BY players.name`,
		sql.Named("teamA", TeamA),
		sql.Named("teamB", TeamB),
	).Scan(&players).Error

	if err != nil {
		return nil, err
	}

	for i := range players {
		players[i].TotalMatches = players[i].MatchesWon + players[i].MatchesDrawn + players[i].MatchesLost
		players[i].TotalPoints = (players[i].MatchesWon * 3) + players[i].MatchesDrawn
		if players[i].TotalMatches > 0 {
			percentage := float64(players[i].MatchesWon*100) / float64(players[i].TotalMatches)
			players[i].VictoryPercentage = math.Round(percentage*100) / 100
		}
	}

	return players, nil
}

func (a *DbAdapter) DeletePlayer(id int) error {
	ctx := context.Background()
	_, err := gorm.G[Player](a.db).Where("id = ?", id).Delete(ctx)

	return err
}

// Seasons
func (a *DbAdapter) CreateSeason(season *Season) error {
	ctx := context.Background()
	return gorm.G[Season](a.db).Create(ctx, season)
}

func (a *DbAdapter) GetSeasons() ([]Season, error) {
	ctx := context.Background()
	return gorm.G[Season](a.db).Find(ctx)
}

func (a *DbAdapter) GetLastSeason() (Season, error) {
	ctx := context.Background()
	return gorm.G[Season](a.db).Last(ctx)
}

// Matches
func (a *DbAdapter) CreateMatch(match *Match) error {
	ctx := context.Background()
	result := gorm.WithResult()
	return gorm.G[Match](a.db, result).Create(ctx, match)
}

func (a *DbAdapter) CreatePlayerMatches(playerMatches *[]PlayerMatch) error {
	ctx := context.Background()
	return gorm.G[PlayerMatch](a.db).CreateInBatches(ctx, playerMatches, 100) // Hardcoded on purpose. There is no need for a configurable value
}
