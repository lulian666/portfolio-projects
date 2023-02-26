package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"portfolio/models"
	"portfolio/pkg"
	"portfolio/services"
)

type ProjectHandler struct {
	service services.ProjectService
}

func NewProjectHandler(s services.ProjectService) ProjectHandler {
	return ProjectHandler{
		service: s,
	}
}

// CreateProject
// @Summery create project
// @Produce json
// @Param name query string false "project name"
// @Router /api/v1/projects [post]
func (h ProjectHandler) CreateProject(c *gin.Context) {
	response := pkg.NewResponse(c)
	var p = models.Project{}

	err := c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.ServerError, err.Error()))
		return
	}

	err = h.service.CreateProject(&p)
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.ServerError, err.Error()))
		return
	}

	response.ToResponse(p)
}

func (h ProjectHandler) DeleteProject(c *gin.Context) {

}

func (h ProjectHandler) UpdateProject(c *gin.Context) {

}

func (h ProjectHandler) GetProject(c *gin.Context) {

}

func (h ProjectHandler) ListProjects(c *gin.Context) {

}
