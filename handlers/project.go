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
	response := pkg.NewResponse(c)
	id := c.Param("id")
	err := h.service.DeleteProject(pkg.StrTo(id).MustInt())
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.NotFound, err.Error()))
		return
	}
	res := struct {
		Message string `json:"message"`
	}{
		Message: "project deleted",
	}
	response.ToResponse(res)
}

func (h ProjectHandler) UpdateProject(c *gin.Context) {
	response := pkg.NewResponse(c)
	id := c.Param("id")
	var p = models.Project{}

	err := c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.ServerError, err.Error()))
		return
	}

	err = h.service.UpdateProject(pkg.StrTo(id).MustInt(), p)
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.NotFound, err.Error()))
		return
	}

	response.ToResponse(p)
}

func (h ProjectHandler) GetProject(c *gin.Context) {
	response := pkg.NewResponse(c)
	id := c.Param("id")
	p, err := h.service.GetProject(pkg.StrTo(id).MustInt())
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.NotFound, err.Error()))
		return
	}
	response.ToResponse(p)
}

func (h ProjectHandler) ListProjects(c *gin.Context) {
	response := pkg.NewResponse(c)
	var p = models.Project{}

	pager := pkg.Pager{
		Page:     pkg.GetPage(c),
		PageSize: pkg.GetPageSize(c),
	}
	projects, totalRows, err := h.service.ListProject(p, pager)
	if err != nil {
		response.ToErrorResponse(pkg.NewError(pkg.NotFound, err.Error()))
		return
	}
	response.ToResponseList(projects, totalRows)
}
