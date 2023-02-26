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
// @Tags projects
// @Summery create project
// @Produce json
// @Param project body models.SwaggerProject false "project"
// @Success 200 {object} models.SwaggerProject "成功"
// @Failure 400 {object} pkg.Error "请求错误"
// @Failure 500 {object} pkg.Error "内部错误"
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

// DeleteProject
// @Tags projects
// @Summery delete project
// @Produce json
// @Param id path int true "project id"
// @Router /api/v1/projects/{id} [delete]
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

// UpdateProject
// @Tags projects
// @Summery update project
// @Produce json
// @Param id path int true "project id"
// @Param project body models.SwaggerProject false "project"
// @Success 200 {object} models.SwaggerProject "成功"
// @Failure 400 {object} pkg.Error "请求错误"
// @Failure 500 {object} pkg.Error "内部错误"
// @Router /api/v1/projects/{id} [put]
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

// GetProject
// @Tags projects
// @Summery get project
// @Produce json
// @Param id path int true "project id"
// @Router /api/v1/projects/{id} [get]
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

// ListProjects
// @Tags projects
// @Summery list project
// @Produce json
// @Param page query string false "page"
// @Param page_size query string false "page size"
// @Success 200 {array} models.SwaggerProjects "成功"
// @Failure 400 {object} pkg.Error "请求错误"
// @Failure 500 {object} pkg.Error "内部错误"
// @Router /api/v1/projects [get]
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
