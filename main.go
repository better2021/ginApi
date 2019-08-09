package main

import (
	"ginApi/controls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello go",
		})
	})

	v2 := router.Group("/api/v2")
	{
		v2.GET("/music", controls.MusicList)
		v2.POST("/music", controls.MusicCreate)
		v2.PUT("/music/:id", controls.MusicUpdate)
		v2.DELETE("/music/:id", controls.MusicDelete)
	}

	router.Run("localhost:8081")
}
