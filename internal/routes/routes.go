package routes

import (
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/matches", handlers.GetMatches)
}
