package main

import (
	"fmt"

	"github.com/JorgeSaicoski/golang-volley-live-score/internal/services/database"
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

	r.Run(":8081")
}
