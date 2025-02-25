package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PromotionController handles Promotion-related requests
type PromotionController struct {
	DB *gorm.DB
}

// NewPromotionController creates a new instance of PromotionController
func NewPromotionController(db *gorm.DB) *PromotionController {
	return &PromotionController{DB: db}
}

// GetPromotions godoc
// @Summary Get all Promotions
// @Description Get a list of all Promotions
// @Tags Promotions
// @Accept json
// @Produce json
// @Success 200 {array} models.Promotion
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Promotions [get]
func (ctrl *PromotionController) GetPromotions(c *gin.Context) {
	var Promotions []models.Promotion
	if err := ctrl.DB.Find(&Promotions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Promotions)
}

// CreatePromotion godoc
// @Summary Create a new Promotion
// @Description Create a new Promotion with the given information
// @Tags Promotions
// @Accept json
// @Produce json
// @Param Promotion body models.Promotion true "Promotion data"
// @Success 201 {object} models.Promotion
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Promotions [post]
func (ctrl *PromotionController) CreatePromotion(c *gin.Context) {
	var Promotion models.Promotion
	if err := c.ShouldBindJSON(&Promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&Promotion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Promotion)
}

// UpdatePromotion godoc
// @Summary Update Promotion
// @Description Update Promotion data by ID
// @Tags Promotions
// @Accept json
// @Produce json
// @Param id path int true "Promotion ID"
// @Param Promotion body models.Promotion true "Updated Promotion data"
// @Success 200 {object} models.Promotion
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Promotions/{id} [put]
func (ctrl *PromotionController) UpdatePromotion(c *gin.Context) {
	id := c.Param("id")
	var Promotion models.Promotion
	if err := ctrl.DB.First(&Promotion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
		return
	}
	if err := c.ShouldBindJSON(&Promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Save(&Promotion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Promotion)
}

// DeletePromotion godoc
// @Summary Delete Promotion
// @Description Delete a Promotion by ID
// @Tags Promotions
// @Accept json
// @Produce json
// @Param id path int true "Promotion ID"
// @Success 204 "No Content"
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Promotions/{id} [delete]
func (ctrl *PromotionController) DeletePromotion(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.Promotion{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
