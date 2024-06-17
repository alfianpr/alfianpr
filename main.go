package main

import (
	"alfianpr/controllers"
	"alfianpr/database"
	"alfianpr/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connectionString := "./database/sqlite/alfianpr.db"
	if err := database.Connect(connectionString); err != nil {
		panic(err)
	}

	router.GET("/api/blog", controllers.GetPosts)
	router.POST("/api/blog", controllers.CreatePost)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		var posts []models.BlogPost
		database.DB.Order("created_at desc").Limit(5).Find(&posts)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Posts": posts,
		})
	})

	router.GET("/api/top5", func(c *gin.Context) {
		var posts []models.BlogPost
		database.DB.Order("created_at desc").Limit(5).Find(&posts)
		c.JSON(http.StatusOK, posts)
	})

	router.Run(":8080")
}
