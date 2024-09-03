package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Set struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	MatchID    uint `json:"match_id"`
	ScoreTeamA int  `json:"score_team_a"`
	ScoreTeamB int  `json:"score_team_b"`
	Win        bool `json:"win"`
}

type Match struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Sets      []Set     `json:"sets" gorm:"foreignKey:MatchID"`
	IsLive    bool      `json:"is_live"`
	Win       bool      `json:"win"`
	MatchDate time.Time `json:"match_date"`
}

func Connect() {
	DB, err := gorm.Open(sqlite.Open("volley_live_data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(&Match{}, &Set{}); err != nil {
		panic("failed to migrate")
	}

	DB.AutoMigrate(&Match{}, &Set{})

}
