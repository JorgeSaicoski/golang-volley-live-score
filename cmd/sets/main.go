package main

import (
	"fmt"

	"github.com/JorgeSaicoski/golang-volley-live-score/internal/api/handlers"
	"github.com/JorgeSaicoski/golang-volley-live-score/internal/services/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Connect(); err != nil {
		fmt.Println("DB not Connected:", err)
		return
	} else {
		fmt.Println("DB Connected")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.GET("/ws", handlers.HandleWebSocker)

	r.Run(":8081")
}
