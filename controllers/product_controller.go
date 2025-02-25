package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductController handles Product-related requests
type ProductController struct {
	DB *gorm.DB
}

// NewProductController creates a new instance of ProductController
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

// GetProducts godoc
// @Summary Get all Products
// @Description Get a list of all Products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Products [get]
func (ctrl *ProductController) GetProducts(c *gin.Context) {
	var Products []models.Product
	if err := ctrl.DB.Find(&Products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Products)
}

// GetProductByID godoc
// @Summary Get Product by ID
// @Description Retrieve a Product by their ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.ErrorResponse
// @Router /api/Products/{id} [get]
func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var Product models.Product
	if err := ctrl.DB.First(&Product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, Product)
}

// CreateProduct godoc
// @Summary Create a new Product
// @Description Create a new Product with the given information
// @Tags Products
// @Accept json
// @Produce json
// @Param Product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Products [post]
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var Product models.Product
	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&Product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

// UpdateProduct godoc
// @Summary Update Product
// @Description Update Product data by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param Product body models.Product true "Updated Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Products/{id} [put]
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var Product models.Product
	if err := ctrl.DB.First(&Product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&Product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Save(&Product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Product)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete a Product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Products/{id} [delete]
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
