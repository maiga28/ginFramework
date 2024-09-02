package controllers

import (
	"ginFramework/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"
)

func GetComments(c *gin.Context) {
	var comments []models.Comment
	if err := db.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := db.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := db.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Comment{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetPostComments(c *gin.Context) {
	postId := c.Param("post_id")
	var comments []models.Comment
	if err := db.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}
