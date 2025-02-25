package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TrainingController обрабатывает запросы, связанные с тренировками.
type TrainingController struct {
	DB *gorm.DB
}

// NewTrainingController создаёт новый экземпляр TrainingController.
func NewTrainingController(db *gorm.DB) *TrainingController {
	return &TrainingController{DB: db}
}

// GetTrainingSessions godoc
// @Summary Получить расписание тренировок
// @Tags Trainings
// @Produce json
// @Success 200 {array} models.TrainingSession
// @Failure 500 {object} map[string]string
// @Router /api/training-sessions [get]
func (ctrl *TrainingController) GetTrainingSessions(c *gin.Context) {
	var sessions []models.TrainingSession
	if err := ctrl.DB.Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

// RegisterForTraining godoc
// @Summary Записаться на тренировку
// @Tags Trainings
// @Accept json
// @Produce json
// @Param payload body models.TrainingRegistration true "Данные регистрации на тренировку"
// @Success 201 {object} models.TrainingRegistration
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/training-registrations [post]
func (ctrl *TrainingController) RegisterForTraining(c *gin.Context) {
	var registration models.TrainingRegistration
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, registration)
}

// GetTrainingParticipants godoc
// @Summary Получить участников тренировки по ID сессии
// @Tags Trainings
// @Produce json
// @Param sessionId path int true "ID тренировки"
// @Success 200 {array} models.TrainingRegistration
// @Failure 500 {object} map[string]string
// @Router /api/training-sessions/{sessionId}/participants [get]
func (ctrl *TrainingController) GetTrainingParticipants(c *gin.Context) {
	sessionID := c.Param("sessionId")
	var registrations []models.TrainingRegistration
	if err := ctrl.DB.Where("training_session_id = ?", sessionID).Find(&registrations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, registrations)
}

// CancelMyTraining godoc
// @Summary Отменить запись на тренировку (в таблице MyTraining)
// @Tags Trainings
// @Produce json
// @Param id path int true "ID записи в MyTraining"
// @Success 200 {object} models.MyTraining
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/my-trainings/{id}/cancel [post]
func (ctrl *TrainingController) CancelMyTraining(c *gin.Context) {
	id := c.Param("id")
	var myTraining models.MyTraining
	if err := ctrl.DB.First(&myTraining, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	myTraining.Status = "отменена"
	if err := ctrl.DB.Save(&myTraining).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, myTraining)
}
