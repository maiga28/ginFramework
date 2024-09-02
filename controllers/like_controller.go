package controllers

import (
	"ginFramework/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"
)

func GetLikes(c *gin.Context) {
	var likes []models.Like
	if err := db.Find(&likes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}

func LikePost(c *gin.Context) {
	var like models.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, like)
}

func UnlikePost(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Like{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetPostLikes(c *gin.Context) {
	postId := c.Param("post_id")
	var likes []models.Like
	if err := db.Where("post_id = ?", postId).Find(&likes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}
