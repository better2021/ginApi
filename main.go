package main

import (
	"ginApi/controls"
	_ "ginApi/docs" // 注意这个一定要引入自己的docs
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Golang Gin API
// @version 2.0
// @description An example of gin
// @termsOfService
// @license.name MIT
func main() {
	/*
		cors.Default() 默认允许所有源跨域
		跨域要写在路由前面，需要先执行
	*/
	router := gin.Default()
	router.Use(cors.Default())

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v2 := router.Group("/api/v2")
	{
		v2.GET("/music", controls.MusicList)
		v2.POST("/music", controls.MusicCreate)
		v2.PUT("/music/:id", controls.MusicUpdate)
		v2.DELETE("/music/:id", controls.MusicDelete)
	}

	router.Run("localhost:8081")
}
