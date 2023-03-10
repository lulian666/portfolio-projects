package routers

import (
	"github.com/gin-gonic/gin"
	_ "portfolio/docs"
	"portfolio/handlers"
	"portfolio/initializers"
	"portfolio/middleware"
	"portfolio/services"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	pService := services.NewProjectService(initializers.DB)
	pHandler := handlers.NewProjectHandler(pService)

	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/projects", middleware.Validator(&middleware.CreateProject{}), pHandler.CreateProject)
		apiV1.DELETE("/projects/:id", pHandler.DeleteProject)
		apiV1.PUT("/projects/:id", middleware.Validator(&middleware.UpdateProject{}), pHandler.UpdateProject)
		apiV1.GET("/projects/:id", pHandler.GetProject)
		apiV1.GET("/projects", pHandler.ListProjects)
	}

	return r
}
