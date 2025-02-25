package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewsController handles News-related requests
type NewsController struct {
	DB *gorm.DB
}

// NewNewsController creates a new instance of NewsController
func NewNewsController(db *gorm.DB) *NewsController {
	return &NewsController{DB: db}
}

// GetNews godoc
// @Summary Get all News
// @Description Get a list of all News
// @Tags News
// @Accept json
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {object} models.ErrorResponse
// @Router /api/News [get]
func (ctrl *NewsController) GetNews(c *gin.Context) {
	var News []models.News
	if err := ctrl.DB.Find(&News).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, News)
}

// CreateNews godoc
// @Summary Create a new News
// @Description Create a new News with the given information
// @Tags News
// @Accept json
// @Produce json
// @Param News body models.News true "News data"
// @Success 201 {object} models.News
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/News [post]
func (ctrl *NewsController) CreateNews(c *gin.Context) {
	var News models.News
	if err := c.ShouldBindJSON(&News); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&News).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, News)
}

// UpdateNews godoc
// @Summary Update News
// @Description Update News data by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Param News body models.News true "Updated News data"
// @Success 200 {object} models.News
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/News/{id} [put]
func (ctrl *NewsController) UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var News models.News
	if err := ctrl.DB.First(&News, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	if err := c.ShouldBindJSON(&News); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Save(&News).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, News)
}

// DeleteNews godoc
// @Summary Delete News
// @Description Delete a News by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 204 "No Content"
// @Failure 500 {object} models.ErrorResponse
// @Router /api/News/{id} [delete]
func (ctrl *NewsController) DeleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.News{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
