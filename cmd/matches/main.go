package main

import (
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/routes"
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/services/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8000")

}
