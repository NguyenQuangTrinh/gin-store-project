package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `gorm:"column:name"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	Age       string    `gorm:"column:age"`
	Role      Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	RoleID    uint      `gorm:"not null;DEFAULT:3" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (user *User) Save() (*User, error) {
	err := Database.Model(&user).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return user, nil
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	err := Database.Where("username=?", username).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func FetchAllUsers() (*[]User, error) {
	var user []User
	err := Database.Find(&user).Error
	if err != nil {
		return &[]User{}, err
	}
	return &user, nil
}

func FetchUser(id string) (*User, error) {
	var user User
	err := Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (user *User) UpdateUser(id string) (*User, error) {
	err := Database.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	err := Database.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
