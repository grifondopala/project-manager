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

func CreateProject(uId uint) (Project, error) {

	project := Project{
		Name:        "Untitled",
		Description: "No description",
		IconSrc:     "/static/project-icons/planning.png",
		Color:       "#94e6e3",
		UserID:      uId,
	}

	var err error
	err = DB.Create(&project).Error
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

func GetUserProjects(uId uint) ([]Project, error) {

	var projects []Project

	if err := DB.Order("updated_at DESC").Find(&projects, "user_id = ?", uId).Error; err != nil {
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

type UpdateProjectInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IconSrc     string `json:"icon_src" binding:"required"`
	Color       string `json:"color" binding:"required"`
	Id          uint   `json:"id" binding:"required"`
}

func UpdateInformation(input UpdateProjectInput) (Project, error) {

	if err := DB.Model(&Project{}).Update(input).Error; err != nil {
		return Project{}, errors.New("project not found")
	}

	return Project{}, nil

}
