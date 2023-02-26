package models

import (
	"gorm.io/gorm"
	"portfolio/initializers"
)

type Project struct {
	gorm.Model
	Name        string
	Description string
	SourceCode  string
	Link        string
}

func (p Project) Create() error {
	result := initializers.DB.Create(&p)
	return result.Error
}

func (p Project) List(pageOffset, pageSize int) ([]Project, error) {
	var projects []Project
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		initializers.DB.Offset(pageOffset).Limit(pageSize)
	}

	if p.Name != "" {
		initializers.DB.Where("name = ?", p.Name)
	}
	if err = initializers.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, err
}

func (p Project) Update() error {
	result := initializers.DB.Model(&Project{}).Where("id = ?", p.ID, 0).Updates(p)
	return result.Error
}

func (p Project) Get() error {
	result := initializers.DB.First(&p)
	return result.Error
}

func (p Project) Delete() error {
	result := initializers.DB.Delete(&p)
	return result.Error
}
