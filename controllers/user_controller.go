package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController handles user-related requests
type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new instance of UserController
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := ctrl.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the given information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
