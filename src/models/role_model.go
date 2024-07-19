package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

func CreateRole(Role *Role) (err error) {
	err = Database.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRoles(Role *[]Role) (err error) {
	err = Database.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRole(Role *Role, id int) (err error) {
	err = Database.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRole(Role *Role) (err error) {
	Database.Save(Role)
	return nil
}
