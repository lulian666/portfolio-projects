package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"portfolio/initializers"
	"portfolio/models"
	"portfolio/pkg"
)

type ProjectHandler struct {
}

func NewProjectHandler() ProjectHandler {
	return ProjectHandler{}
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

	err = initializers.DB.Create(&p).Error
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
