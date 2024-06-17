package main

import (
	"alfianpr/controllers"
	"alfianpr/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connectionString := "/database/sqlite/alfianpr.db"
	if err := database.Connect(connectionString); err != nil {
		panic(err)
	}

	router.GET("/api/blog", controllers.GetPosts)
	router.POST("/api/blog", controllers.CreatePost)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":8080")
}
