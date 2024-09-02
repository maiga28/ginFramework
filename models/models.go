package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post"`
}

type Like struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `json:"user"`
	PostID uint `json:"post_id"`
	Post   Post `json:"post"`
}
