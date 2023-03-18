package models

import (
	"errors"
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

func (t *TextPoint) UpdateTextPoint() (*TextPoint, error) {

	if err := DB.Model(&TextPoint{}).Update(t).Error; err != nil {
		return t, errors.New("text point not found")
	}

	return t, nil

}
