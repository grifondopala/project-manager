package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	Description string `gorm:"size:255;not null;" json:"description"`
	IconSrc     string `gorm:"size:255;not null;" json:"icon_src"`
	Color       string `gorm:"size:255;not null;" json:"color"`
	UserID      uint   `gorm:"not null;" json:"user_id"`
}

func (p Project) SaveProject() (Project, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return Project{}, err
	}
	return p, nil
}

func GetUserProjects(uId uint) ([]Project, error) {

	var projects []Project

	if err := DB.Find(&projects, "user_id = ?", uId).Error; err != nil {
		return projects, errors.New("projects not found")
	}

	return projects, nil

}

func GetProjectById(pId uint) (Project, error) {

	var project Project

	if err := DB.First(&project, pId).Error; err != nil {
		return project, errors.New("project not found")
	}

	return project, nil

}
