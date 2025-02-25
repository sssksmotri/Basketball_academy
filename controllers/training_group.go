package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TrainingGroupController handles TrainingGroup-related requests
type TrainingGroupController struct {
	DB *gorm.DB
}

// NewTrainingGroupController creates a new instance of TrainingGroupController
func NewTrainingGroupController(db *gorm.DB) *TrainingGroupController {
	return &TrainingGroupController{DB: db}
}

// GetTrainingGroup godoc
// @Summary Get all TrainingGroup
// @Description Get a list of all TrainingGroup
// @Tags TrainingGroup
// @Accept json
// @Produce json
// @Success 200 {array} models.TrainingGroup
// @Failure 500 {object} models.ErrorResponse
// @Router /api/TrainingGroup [get]
func (ctrl *TrainingGroupController) GetTrainingGroup(c *gin.Context) {
	var TrainingGroup []models.TrainingGroup
	if err := ctrl.DB.Find(&TrainingGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, TrainingGroup)
}

// CreateTrainingGroup godoc
// @Summary Create a new TrainingGroup
// @Description Create a new TrainingGroup with the given information
// @Tags TrainingGroup
// @Accept json
// @Produce json
// @Param TrainingGroup body models.TrainingGroup true "TrainingGroup data"
// @Success 201 {object} models.TrainingGroup
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/TrainingGroup [post]
func (ctrl *TrainingGroupController) CreateTrainingGroup(c *gin.Context) {
	var TrainingGroup models.TrainingGroup
	if err := c.ShouldBindJSON(&TrainingGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&TrainingGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, TrainingGroup)
}
