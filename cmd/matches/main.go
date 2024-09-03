package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("volley_live_data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	r.GET("/matches", func(c *gin.Context) {
		var matches []Match
		db.Preload("Sets").Find(&matches)
		c.JSON(200, matches)
	})

	r.Run(":8000")

}
