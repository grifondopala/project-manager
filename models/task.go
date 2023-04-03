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
	Done        bool `gorm:"not null;default:false" json:"done"`
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

func (t *Task) Update() (task *Task, err error) {
	if err := DB.Model(&t).Updates(map[string]interface{}{"order_number": t.OrderNumber, "section_id": t.SectionID, "done": t.Done}).Error; err != nil {
		return &Task{}, errors.New("task not found")
	}
	return t, nil
}
