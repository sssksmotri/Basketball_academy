package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ChatController handles Chat-related requests
type ChatController struct {
	DB *gorm.DB
}

// NewChatController creates a new instance of ChatController
func NewChatController(db *gorm.DB) *ChatController {
	return &ChatController{DB: db}
}

// GetChat godoc
// @Summary Get all Chat
// @Description Get a list of all Chat
// @Tags Chat
// @Accept json
// @Produce json
// @Success 200 {array} models.Chat
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Chat [get]
func (ctrl *ChatController) GetChat(c *gin.Context) {
	var Chat []models.Chat
	if err := ctrl.DB.Find(&Chat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Chat)
}

// CreateChat godoc
// @Summary Create a new Chat
// @Description Create a new Chat with the given information
// @Tags Chat
// @Accept json
// @Produce json
// @Param Chat body models.Chat true "Chat data"
// @Success 201 {object} models.Chat
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/Chat [post]
func (ctrl *ChatController) CreateChat(c *gin.Context) {
	var Chat models.Chat
	if err := c.ShouldBindJSON(&Chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&Chat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Chat)
}
