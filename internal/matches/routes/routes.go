package routes

import (
	handlers_matches "github.com/JorgeSaicoski/golang-volley-live-score/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/matches", handlers_matches.GetMatches)
	r.GET("/matches/:id", handlers_matches.GetMatchByID)
	r.POST("/matches", handlers_matches.CreateMatch)
	r.PUT("/matches/:id", handlers_matches.UpdateMatch)
}
