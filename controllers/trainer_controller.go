package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TrainerController обрабатывает запросы, связанные с тренерами.
type TrainerController struct {
	DB *gorm.DB
}

// NewTrainerController создаёт новый экземпляр TrainerController.
func NewTrainerController(db *gorm.DB) *TrainerController {
	return &TrainerController{DB: db}
}

// GetAllTrainers godoc
// @Summary Получить всех тренеров
// @Tags Trainers
// @Produce json
// @Success 200 {array} models.Trainer
// @Failure 500 {object} map[string]string
// @Router /api/trainers [get]
func (ctrl *TrainerController) GetAllTrainers(c *gin.Context) {
	var trainers []models.Trainer
	if err := ctrl.DB.Find(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trainers)
}

// GetTrainerByID godoc
// @Summary Получить тренера по ID
// @Tags Trainers
// @Produce json
// @Param id path int true "ID тренера"
// @Success 200 {object} models.Trainer
// @Failure 404 {object} map[string]string
// @Router /api/trainers/{id} [get]
func (ctrl *TrainerController) GetTrainerByID(c *gin.Context) {
	id := c.Param("id")
	var trainer models.Trainer
	if err := ctrl.DB.First(&trainer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тренер не найден"})
		return
	}
	c.JSON(http.StatusOK, trainer)
}

// UpdateTrainingSession godoc
// @Summary Обновить данные тренировки (только для тренера)
// @Tags Trainers
// @Accept json
// @Produce json
// @Param id path int true "ID тренировки"
// @Param payload body models.TrainingSession true "Новые данные тренировки"
// @Success 200 {object} models.TrainingSession
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/training-sessions/{id} [put]
func (ctrl *TrainerController) UpdateTrainingSession(c *gin.Context) {
	id := c.Param("id")
	var session models.TrainingSession
	if err := ctrl.DB.First(&session, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тренировка не найдена"})
		return
	}
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Save(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}
