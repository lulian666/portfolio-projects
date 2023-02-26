package routers

import (
	"github.com/gin-gonic/gin"
	_ "portfolio/docs"
	"portfolio/handlers"
	"portfolio/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/projects", middleware.Validator(&middleware.CreateProject{}), handlers.CreateProject)
		apiV1.DELETE("/projects/:id", handlers.DeleteProject)
		apiV1.PUT("/projects/:id", handlers.UpdateProject)
		apiV1.PATCH("/projects/:id/state", handlers.UpdateProject)
		apiV1.GET("/projects/:id", handlers.GetProject)
		apiV1.GET("/projects", handlers.ListProjects)
	}

	return r
}
