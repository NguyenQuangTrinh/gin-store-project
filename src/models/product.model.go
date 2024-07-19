package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"store/src/DTO"
	"strconv"
	"time"
)

type Image struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"productId"`
	Base64    string    `gorm:"type:text" json:"base64"`
}

type Product struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
	NameProduct string    `gorm:"type:char(30);" json:"nameProduct"`
	Description string    `gorm:"type:text;" json:"description"`
	Price       float64   `gorm:"type:decimal;" json:"price"`
	Quantity    int       `gorm:"type:int;" json:"quantity"`
	IsActive    bool      `gorm:"default:true;" json:"isActive"`
	CategoryID  uuid.UUID `gorm:"type:uuid" json:"categoryId"`
	Images      []Image   `gorm:"foreignKey:ProductID;" json:"images"`
	Comments    []Comment `gorm:"foreignKey:ProductID;" json:"comments"`
	CreatedAt   time.Time `gorm:"type:timestamp;" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"type:timestamp;" json:"updatedAt"`
}

type NewProductRequest struct {
	NameProduct string                  `json:"nameProduct"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Quantity    int                     `json:"quantity"`
	IsActive    bool                    `json:"isActive"`
	CategoryID  uuid.UUID               `json:"categoryId"`
	Images      []*multipart.FileHeader `json:"images"`
}

func SaveProduct(req *gin.Context) (*Product, error) {

	price, err := strconv.ParseFloat(req.PostForm("price"), 64)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse: %v", err)
	}
	quantity, errs := strconv.Atoi(req.PostForm("quantity"))
	if errs != nil {
		return nil, fmt.Errorf("Failed to parse: %v", errs)
	}
	categoryId, errCategory := uuid.Parse(req.PostForm("categoryId"))
	if errCategory != nil {
		return nil, fmt.Errorf("Failed to parse: %v", errCategory)
	}

	form, err := req.MultipartForm()
	if err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form"})
		//return
	}

	files := form.File["images"]
	if files == nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": "No images uploaded"})
		//return
	}

	newProduct := Product{
		ID:          uuid.New(),
		NameProduct: req.PostForm("nameProduct"),
		Description: req.PostForm("description"),
		Price:       price,
		Quantity:    quantity,
		IsActive:    true,
		CategoryID:  categoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := Database.Create(&newProduct)

	if result.Error != nil {
		return &Product{}, result.Error
	}

	productId := newProduct.ID
	for _, files := range files {
		fmt.Println("Processing image...")

		// Improved error handling with detailed messages
		file, err := files.Open()
		if err != nil {
			return &Product{}, fmt.Errorf("error opening image file: %w", err)
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			return &Product{}, fmt.Errorf("error copying image data: %w", err)
		}

		// Optional image validation (e.g., size, format)
		// You can use libraries like "github.com/nfnt/resize" for resizing
		// and "github.com/disintegration/imaging" for format checks.

		base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())

		image := Image{
			ID:        uuid.New(),
			ProductID: productId,
			Base64:    base64Image,
		}

		// Create a transaction to ensure data consistency
		err = Database.Transaction(func(tx *gorm.DB) error {

			// Create the image within the transaction using the product ID
			if err := tx.Create(&image).Error; err != nil {
				return fmt.Errorf("error creating image: %w", err)
			}

			return nil // Commit the transaction if both creations succeed
		})

		if err != nil {
			return &Product{}, err
		}
	}

	var resultProduct Product

	if err := Database.Preload("Images").Preload("Comments").Where("id = ?", newProduct.ID).First(&resultProduct).Error; err != nil {
		return nil, fmt.Errorf("failed to find product: %w", err)
	}

	return &resultProduct, nil
}

func GetProductById(productId string) (*DTO.ProductWithComments, error) {
	var product Product

	if err := Database.Preload("Images").Where("id", productId).Find(&product).Error; err != nil {
		return &DTO.ProductWithComments{}, err
	}

	var commentWithUsers []DTO.CommentWithUser

	if err := Database.Table("comments").
		Select("comments.id, comments.content, comments.rating, users.username, users.name").
		Joins("join users on comments.user_id = users.id").
		Where("comments.product_id = ?", productId).
		Order("comments.created_at").
		Limit(10).
		Scan(&commentWithUsers).Error; err != nil {

		return &DTO.ProductWithComments{}, err
	}

	var images []DTO.ImageDTO
	for _, img := range product.Images {
		images = append(images, DTO.ImageDTO{
			ID:        img.ID,
			ProductID: img.ProductID,
			Base64:    img.Base64,
		})
	}

	productWithComments := DTO.ProductWithComments{
		ID:          product.ID,
		NameProduct: product.NameProduct,
		Description: product.Description,
		Price:       product.Price,
		Images:      images,
		Quantity:    product.Quantity,
		IsActive:    product.IsActive,
		Comments:    commentWithUsers,
	}

	return &productWithComments, nil
}

func GetProduct(limit int, page int) (*[]Product, error) {

	var products []Product

	offset := (page - 1) * limit
	if err := Database.Preload("Images").Order("created_at desc").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return &[]Product{}, err
	}

	return &products, nil
}
