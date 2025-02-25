package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PaymentController обрабатывает запросы, связанные с платежами.
type PaymentController struct {
	DB *gorm.DB
}

// NewPaymentController создаёт новый экземпляр PaymentController.
func NewPaymentController(db *gorm.DB) *PaymentController {
	return &PaymentController{DB: db}
}

// GetUserPayments godoc
// @Summary Получить историю платежей пользователя
// @Tags Payments
// @Produce json
// @Param userId path int true "ID пользователя"
// @Success 200 {array} models.Payment
// @Failure 500 {object} map[string]string
// @Router /api/payments/user/{userId} [get]
func (ctrl *PaymentController) GetUserPayments(c *gin.Context) {
	userID := c.Param("userId")
	var payments []models.Payment
	if err := ctrl.DB.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

// GetPaymentDetails godoc
// @Summary Получить детали платежа по ID
// @Tags Payments
// @Produce json
// @Param id path int true "ID платежа"
// @Success 200 {object} models.Payment
// @Failure 404 {object} map[string]string
// @Router /api/payments/{id} [get]
func (ctrl *PaymentController) GetPaymentDetails(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	if err := ctrl.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Платеж не найден"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
