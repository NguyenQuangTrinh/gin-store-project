package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
	Content   string    `gorm:"type:text;" json:"content"`
	Rating    int       `gorm:"type:int;" json:"rating"`
	ProductID uuid.UUID `gorm:"type:uuid;" json:"productId"`
	Product   Product   `gorm:"foreignKey:ProductID;" json:"product"`
	UserID    uuid.UUID `gorm:"type:uuid;" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;" json:"user"`
	CreatedAt time.Time `gorm:"type:timestamp;" json:"createdAt"`
}

type ResultComments struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uuid.UUID `json:"userId"` // UserID để lấy thông tin người dùng
	Username  string    `json:"username"`
	Name      string    `json:"name"`
}

func (comment *Comment) SaveComment() (*ResultComments, error) {
	err := Database.Model(&comment).Save(&comment).Error
	if err != nil {
		return &ResultComments{}, err
	}

	var result ResultComments
	err = Database.Table("comments").
		Select("comments.id, comments.content, comments.rating, comments.created_at, comments.user_id, users.username, users.name").
		Joins("JOIN users ON comments.user_id = users.id").
		Where("comments.id = ?", comment.ID).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func GetCommentByIdProduct(productId string) (*[]ResultComments, error) {
	var resultComments []ResultComments

	var err = Database.Table("comments").
		Select("comments.id, comments.content, comments.rating, comments.product_id, users.id AS user_id, users.username, users.name").
		Joins("JOIN users ON comments.user_id = users.id").
		Where("comments.product_id =?", productId).
		Scan(&resultComments).Error
	if err != nil {
		return &[]ResultComments{}, err
	}

	return &resultComments, nil
}

func DeleteComment(commentId string, userid string) error {
	result := Database.Model(&Comment{}).Where("id=? AND user_id=?", commentId, userid).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("comment does not exist or no p")
	}

	return nil
}

func DeleteCommentAdmin(commentId string) error {
	result := Database.Model(&Comment{}).Where("id=?", commentId).Delete(&Comment{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("comment does not exist or no p")
	}

	return nil
}
