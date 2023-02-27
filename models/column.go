package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Column struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	OrderNumber uint   `gorm:"not null;" json:"order_number"`
	Type        string `gorm:"not null;" json:"type"`
	Width       uint   `gorm:"not null;" json:"width"`
	ProjectID   uint   `gorm:"not null;" json:"project_id"`
}

func (c *Column) Create() (*Column, error) {

	var err error
	err = DB.Create(&c).Error
	if err != nil {
		return &Column{}, err
	}
	return c, nil

}

func (c *Column) GetProjectColumns(pId uint) (columns []*Column, err error) {
	if err := DB.Order("order_number ASC").Find(&columns, "project_id = ?", pId).Error; err != nil {
		return nil, errors.New("columns not found")
	}
	return columns, nil
}

type UpdateColumInput struct {
	Name        string `json:"name" binding:"required"`
	Id          uint   `json:"id" binding:"required"`
	Width       uint   `gorm:"not null;" json:"width"`
	OrderNumber uint   `gorm:"not null;" json:"order_number"`
}

func UpdateColumn(input UpdateColumInput) (Column, error) {

	if err := DB.Model(&Column{}).Update(input).Error; err != nil {
		return Column{}, errors.New("column not found")
	}

	return Column{}, nil

}
