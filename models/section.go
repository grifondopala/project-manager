package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Section struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	OrderNumber uint   `gorm:"not null;" json:"order_number"`
	ProjectID   uint   `gorm:"not null;" json:"project_id"`
}

func (s *Section) Create() (*Section, error) {

	var err error
	err = DB.Create(&s).Error
	if err != nil {
		return &Section{}, err
	}
	return s, nil

}

func (s *Section) GetProjectSections(pId uint) (sections []*Section, err error) {
	if err := DB.Order("order_number ASC").Find(&sections, "project_id = ?", pId).Error; err != nil {
		return nil, errors.New("sections not found")
	}
	return sections, nil
}

func UpdateSection(section Section) (Section, error) {

	if err := DB.Model(&Section{}).Update(section).Error; err != nil {
		return Section{}, errors.New("section not found")
	}

	return Section{}, nil

}
