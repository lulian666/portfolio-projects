package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"portfolio/initializers"
	"portfolio/models"
	"portfolio/pkg"
)

type Project struct {
}

// CreateProject
// @Summery 创建项目
// @Produce json
// @Param name query string false "project name"
// @Router /api/v1/projects [post]
func CreateProject(c *gin.Context) {
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

func DeleteProject(c *gin.Context) {

}

func UpdateProject(c *gin.Context) {

}

func GetProject(c *gin.Context) {

}

func ListProjects(c *gin.Context) {

}
