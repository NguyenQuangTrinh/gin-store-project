package models

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(100);not null" json:"name"`
}

func (category *Category) SaveCategory() (*Category, error) {
	err := Database.Model(&category).Save(&category).Error
	if err != nil {
		return &Category{}, err
	}

	return category, nil
}

func UpdateCategory(category Category) (*Category, error) {
	err := Database.Model(&Category{}).Where("id = ?", category.ID).Updates(&category).Error
	if err != nil {
		return &Category{}, err
	}

	return &category, nil
}
