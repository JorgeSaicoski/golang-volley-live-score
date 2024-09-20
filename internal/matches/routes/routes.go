package routes

import (
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/matches/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/matches", handlers.GetMatches)
	r.GET("/matches/:id", handlers.GetMatchByID)
	r.POST("/matches", handlers.CreateMatch)
	r.PUT("/matches/:id", handlers.UpdateMatch)
}
