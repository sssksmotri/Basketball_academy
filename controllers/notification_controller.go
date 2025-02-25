package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NotificationController обрабатывает запросы, связанные с уведомлениями.
type NotificationController struct {
	DB *gorm.DB
}

// NewNotificationController создаёт новый экземпляр NotificationController.
func NewNotificationController(db *gorm.DB) *NotificationController {
	return &NotificationController{DB: db}
}

// SendNotification godoc
// @Summary Отправить уведомление
// @Tags Notifications
// @Accept json
// @Produce json
// @Param payload body models.Notification true "Данные уведомления"
// @Success 201 {object} models.Notification
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/notifications [post]
func (ctrl *NotificationController) SendNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, notification)
}
