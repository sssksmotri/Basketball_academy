package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderController handles Order-related requests
type OrderController struct {
	DB *gorm.DB
}

// NewOrderController creates a new instance of OrderController
func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

// Getorders godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Router /api/orders [get]
func (ctrl *OrderController) GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := ctrl.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CreateOrder godoc
// @Summary Create a new Order
// @Description Create a new Order with the given information
// @Tags orders
// @Accept json
// @Produce json
// @Param Order body models.Order true "Order data"
// @Success 201 {object} models.Order
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var Order models.Order
	if err := c.ShouldBindJSON(&Order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&Order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Order)
}
