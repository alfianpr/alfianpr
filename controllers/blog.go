package controllers

import (
	"alfianpr/database"
	"alfianpr/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func GetPosts(c *gin.Context) {
	var posts []models.BlogPost
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.BlogPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.CreatedAt = time.Now()
	post.HTMLContent = string(blackfriday.Run([]byte(post.Content)))
	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}
