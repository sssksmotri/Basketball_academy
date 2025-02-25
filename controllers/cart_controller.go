package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CartController обрабатывает запросы, связанные с корзиной.
type CartController struct {
	DB *gorm.DB
}

// NewCartController создаёт новый экземпляр CartController.
func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

// GetCartByUser godoc
// @Summary Получить корзину пользователя по ID
// @Tags Cart
// @Produce json
// @Param userId path int true "ID пользователя"
// @Success 200 {object} models.Cart
// @Failure 404 {object} map[string]string
// @Router /api/carts/user/{userId} [get]
func (ctrl *CartController) GetCartByUser(c *gin.Context) {
	userID := c.Param("userId")
	var cart models.Cart
	if err := ctrl.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Корзина не найдена"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// AddToCart godoc
// @Summary Добавить товар в корзину
// @Tags Cart
// @Accept json
// @Produce json
// @Param payload body models.CartItem true "Данные товара в корзине"
// @Success 201 {object} models.CartItem
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart-items [post]
func (ctrl *CartController) AddToCart(c *gin.Context) {
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// Checkout godoc
// @Summary Оформить заказ из корзины
// @Tags Cart
// @Accept json
// @Produce json
// @Param payload body models.Cart true "ID корзины"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/carts/checkout [post]
func (ctrl *CartController) Checkout(c *gin.Context) {
	var data models.Cart
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Логика перевода товаров из корзины в заказ, очистка корзины и т.д.
	c.JSON(http.StatusOK, gin.H{"message": "Заказ оформлен"})
}
