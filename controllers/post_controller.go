package controllers

import (
	"ginFramework/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetUserPosts(c *gin.Context) {
	userId := c.Param("user_id")
	var posts []models.Post
	if err := db.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
