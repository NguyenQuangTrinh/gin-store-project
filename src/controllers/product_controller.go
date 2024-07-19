package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/src/models"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	//var input models.NewProductRequest

	//form, _ := c.MultipartForm()
	//files := form.File["Images"]

	//if err := c.ShouldBindJSON(&input); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": "failed input", "message": err.Error(), "data": nil})
	//	return
	//}
	saveProduct, err := models.SaveProduct(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed result", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": saveProduct})
}

func GetProduct(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 10
	}

	product, err := models.GetProduct(pageSizeInt, pageInt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": product})

}

func GetProductById(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}

	product, err := models.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": product})

}
