package models

import (
	"github.com/jinzhu/gorm"
)

type TextPoint struct {
	gorm.Model
	Text     string `gorm:"not null;" json:"text"`
	TaskID   uint   `gorm:"not null;" json:"task_id"`
	ColumnID uint   `gorm:"not null;" json:"column_id"`
}

func (t *TextPoint) Create() (*TextPoint, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &TextPoint{}, err
	}
	return t, nil

}

func (t *TextPoint) GetTextPoint(tId uint, cId uint) (*TextPoint, error) {

	err := DB.First(&t, "task_id = ? and column_id = ?", tId, cId).Error
	if err != nil {
		return &TextPoint{}, err
	}
	return t, err

}
