package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	_ "portfolio/docs"
	"portfolio/initializers"
	"portfolio/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

// main
// @title personal projects
// @version 1.0
func main() {
	r := routers.NewRouter()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run()
	if err != nil {
		log.Fatal("Failed to start server, due to: ", err.Error())
	}
}
