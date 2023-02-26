package services

import (
	"gorm.io/gorm"
	"portfolio/models"
	"portfolio/pkg"
)

type ProjectService interface {
	CreateProject(param *models.Project) error
	GetProject(ID int) (models.Project, error)
	ListProject(param models.Project, pager pkg.Pager) ([]models.Project, int64, error)
	DeleteProject(id int) error
	UpdateProject(id int, param models.Project) error
}

type projectService struct {
	DB *gorm.DB
}

func NewProjectService(db *gorm.DB) ProjectService {
	return &projectService{
		DB: db,
	}
}

func (s projectService) CreateProject(param *models.Project) error {
	return param.Create(s.DB)
}

func (s projectService) GetProject(id int) (models.Project, error) {
	var p = models.Project{}
	return p.Get(s.DB, id)
}

func (s projectService) ListProject(param models.Project, pager pkg.Pager) ([]models.Project, int64, error) {
	pageOffset := pkg.GetPageOffset(pager.Page, pager.PageSize)
	projects, totalRows, err := param.List(s.DB, pageOffset, pager.PageSize)
	return projects, totalRows, err
}

func (s projectService) DeleteProject(id int) error {
	var p = models.Project{}
	return p.Delete(s.DB, id)
}

func (s projectService) UpdateProject(id int, param models.Project) error {
	return param.Update(s.DB, id)
}
