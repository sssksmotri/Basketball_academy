package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PromoCodeController handles PromoCode-related requests
type PromoCodeController struct {
	DB *gorm.DB
}

// NewPromoCodeController creates a new instance of PromoCodeController
func NewPromoCodeController(db *gorm.DB) *PromoCodeController {
	return &PromoCodeController{DB: db}
}

// GetPromoCodes godoc
// @Summary Get all PromoCodes
// @Description Get a list of all PromoCodes
// @Tags PromoCodes
// @Accept json
// @Produce json
// @Success 200 {array} models.PromoCode
// @Failure 500 {object} models.ErrorResponse
// @Router /api/PromoCodes [get]
func (ctrl *PromoCodeController) GetPromoCodes(c *gin.Context) {
	var PromoCodes []models.PromoCode
	if err := ctrl.DB.Find(&PromoCodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PromoCodes)
}

// CreatePromoCode godoc
// @Summary Create a new PromoCode
// @Description Create a new PromoCode with the given information
// @Tags PromoCodes
// @Accept json
// @Produce json
// @Param PromoCode body models.PromoCode true "PromoCode data"
// @Success 201 {object} models.PromoCode
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/PromoCodes [post]
func (ctrl *PromoCodeController) CreatePromoCode(c *gin.Context) {
	var PromoCode models.PromoCode
	if err := c.ShouldBindJSON(&PromoCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&PromoCode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, PromoCode)
}
