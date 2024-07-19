package DTO

import (
	"github.com/google/uuid"
)

type CommentWithUser struct {
	ID       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	Rating   int       `json:"rating"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
}
type ImageDTO struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productId"`
	Base64    string    `json:"base64"`
}

type ProductWithComments struct {
	ID          uuid.UUID         `json:"id"`
	NameProduct string            `json:"nameProduct"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Quantity    int               `json:"quantity"`
	Images      []ImageDTO        `json:"images"`
	IsActive    bool              `json:"isActive"`
	Comments    []CommentWithUser `json:"comments"`
}
