package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	OrderNumber uint `gorm:"not null;" json:"order_number"`
	ProjectID   uint `gorm:"not null;" json:"project_id"`
	SectionID   uint `gorm:"not null;" json:"section_id"`
}

func (t *Task) Create() (*Task, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &Task{}, err
	}
	return t, nil

}

func (t *Task) GetProjectTasks(pId uint) (tasks []*Task, err error) {
	if err := DB.Order("section_id ASC").Find(&tasks, "project_id = ?", pId).Error; err != nil {
		return nil, errors.New("tasks not found")
	}
	return tasks, nil
}
