package handlers

import (
	"net/http"
	"strconv"

	"github.com/JorgeSaicoski/golang-volley-live-score/internal/services/database"
	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	var matches []database.Match
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset := (page - 1) * pageSize

	if err := database.DB.Preload("Sets").Offset(offset).Limit(pageSize).Find(&matches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, matches)
}

func GetMatchByID(c *gin.Context) {
	var match database.Match

	matchId := c.Param("id")

	if err := database.DB.Preload("Sets").First(&match, matchId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, match)
}

func CreateMatch(c *gin.Context) {
	var match database.Match

	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, match)
}

func UpdateMatch(c *gin.Context) {
	var match database.Match

	matchId := c.Param("id")

	if err := database.DB.Preload("Sets").First(&match, matchId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, match)

}
