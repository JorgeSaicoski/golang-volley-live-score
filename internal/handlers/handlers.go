package handlers

import (
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/services/database"
	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	var matches []database.Match
	database.DB.Preload("Sets").Find(&matches)
	c.JSON(200, matches)
}
