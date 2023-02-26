package services

import (
	"gorm.io/gorm"
	"portfolio/models"
)

type ProjectService interface {
	CreateProject(param *models.Project) error
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
	err := s.DB.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}
