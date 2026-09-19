package db

import (
	"context"
	"log"

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

func (a *DbAdapter) GetPlayers() ([]Player, error) {
	ctx := context.Background()
	return gorm.G[Player](a.db).Find(ctx)
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
