package controllers

import (
	"net/http"

	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ClubController обрабатывает запросы, связанные с клубами.
type ClubController struct {
	DB *gorm.DB
}

// NewClubController создаёт новый экземпляр ClubController.
func NewClubController(db *gorm.DB) *ClubController {
	return &ClubController{DB: db}
}

// GetAllClubs godoc
// @Summary Получить все клубы
// @Tags Clubs
// @Produce json
// @Success 200 {array} models.Club
// @Failure 500 {object} map[string]string
// @Router /api/clubs [get]
func (ctrl *ClubController) GetAllClubs(c *gin.Context) {
	var clubs []models.Club
	if err := ctrl.DB.Find(&clubs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clubs)
}

// GetClubByID godoc
// @Summary Получить клуб по ID
// @Tags Clubs
// @Produce json
// @Param id path int true "ID клуба"
// @Success 200 {object} models.Club
// @Failure 404 {object} map[string]string
// @Router /api/clubs/{id} [get]
func (ctrl *ClubController) GetClubByID(c *gin.Context) {
	id := c.Param("id")
	var club models.Club
	if err := ctrl.DB.First(&club, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Клуб не найден"})
		return
	}
	c.JSON(http.StatusOK, club)
}
