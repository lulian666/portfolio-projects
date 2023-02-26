package models

import (
	"gorm.io/gorm"
	"portfolio/pkg"
)

type SwaggerProject struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"required,min=2,max=500"`
	SourceCode  string `json:"sourceCode" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type SwaggerProjects struct {
	List  []*SwaggerProject
	Pager *pkg.Pager
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	SourceCode  string
	Link        string
}

func (p *Project) Create(db *gorm.DB) error {
	result := db.Create(p)
	return result.Error
}

func (p *Project) List(db *gorm.DB, pageOffset, pageSize int) ([]Project, int64, error) {
	var projects []Project
	var err error

	result := db.Find(&projects)
	if result.Error != nil {
		return nil, 0, err
	}
	totalRows := result.RowsAffected

	if pageOffset >= 0 && pageSize > 0 {
		result = result.Offset(pageOffset).Limit(pageSize).Find(&projects)
	}

	if p.Name != "" {
		result = result.Where("name = ?", p.Name).Find(&projects)
	}

	return projects, totalRows, err
}

func (p *Project) Update(db *gorm.DB, id int) error {
	result := db.Model(&Project{}).Where("id = ?", id).Updates(p)

	return result.Error
}

func (p *Project) Get(db *gorm.DB, id int) (Project, error) {
	result := db.First(&p, id)
	if result.Error != nil {
		return Project{}, result.Error
	}
	return *p, nil
}

func (p *Project) Delete(db *gorm.DB, id int) error {
	result := db.Delete(&p, id)
	return result.Error
}
