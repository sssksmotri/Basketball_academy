package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MyTrainingController обрабатывает запросы, связанные с уведомлениями.
type MyTrainingController struct {
	DB *gorm.DB
}

// NewMyTrainingController создаёт новый экземпляр MyTrainingController.
func NewMyTrainingController(db *gorm.DB) *MyTrainingController {
	return &MyTrainingController{DB: db}
}

// SendMyTraining godoc
// @Summary Отправить уведомление
// @Tags MyTrainings
// @Accept json
// @Produce json
// @Param payload body models.MyTraining true "Данные уведомления"
// @Success 201 {object} models.MyTraining
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/MyTrainings [post]
func (ctrl *MyTrainingController) SendMyTraining(c *gin.Context) {
	var MyTraining models.MyTraining
	if err := c.ShouldBindJSON(&MyTraining); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&MyTraining).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, MyTraining)
}
