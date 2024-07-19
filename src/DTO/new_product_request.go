package DTO

import (
	"github.com/google/uuid"
	"mime/multipart"
)

type NewProductRequest struct {
	NameProduct string                  `json:"nameProduct"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Quantity    int                     `json:"quantity"`
	IsActive    bool                    `json:"isActive"`
	CategoryID  uuid.UUID               `json:"categoryId"`
	Images      []*multipart.FileHeader `json:"images"`
}
