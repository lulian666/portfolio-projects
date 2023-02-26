package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/dvwright/xss-mw"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	_ "portfolio/docs"
	"portfolio/initializers"
	"portfolio/routers"
	"time"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Hour // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // handle exceed rate limit request
	}))

	var xssMdlwr xss.XssMw
	r.Use(xssMdlwr.RemoveXss())

	r.Use(helmet.Default())

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
