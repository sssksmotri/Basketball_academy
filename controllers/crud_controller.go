package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CRUDController реализует базовые операции для работы с моделью T.
type CRUDController[T any] struct {
	DB *gorm.DB
}

// NewCRUDController возвращает новый экземпляр CRUDController.
func NewCRUDController[T any](db *gorm.DB) *CRUDController[T] {
	return &CRUDController[T]{DB: db}
}

// GetAll godoc
// @Summary Получить список записей
// @Description Получить все записи модели
// @Tags generic
// @Accept json
// @Produce json
// @Success 200 {array} interface{}
// @Failure 500 {object} map[string]string
// @Router /{resource} [get]
func (ctrl *CRUDController[T]) GetAll(c *gin.Context) {
	var items []T
	if err := ctrl.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetByID godoc
// @Summary Получить запись по ID
// @Description Получить запись модели по идентификатору
// @Tags generic
// @Accept json
// @Produce json
// @Param id path int true "ID записи"
// @Success 200 {object} interface{}
// @Failure 404 {object} map[string]string
// @Router /{resource}/{id} [get]
func (ctrl *CRUDController[T]) GetByID(c *gin.Context) {
	id := c.Param("id")
	var item T
	if err := ctrl.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Create godoc
// @Summary Создать новую запись
// @Description Создать новую запись модели
// @Tags generic
// @Accept json
// @Produce json
// @Param data body interface{} true "Данные записи"
// @Success 201 {object} interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /{resource} [post]
func (ctrl *CRUDController[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// Update godoc
// @Summary Обновить запись
// @Description Обновить существующую запись по ID
// @Tags generic
// @Accept json
// @Produce json
// @Param id path int true "ID записи"
// @Param data body interface{} true "Данные для обновления"
// @Success 200 {object} interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /{resource}/{id} [put]
func (ctrl *CRUDController[T]) Update(c *gin.Context) {
	id := c.Param("id")
	var item T
	if err := ctrl.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Delete godoc
// @Summary Удалить запись
// @Description Удалить запись модели по ID
// @Tags generic
// @Accept json
// @Produce json
// @Param id path int true "ID записи"
// @Success 204 "No Content"
// @Failure 500 {object} map[string]string
// @Router /{resource}/{id} [delete]
func (ctrl *CRUDController[T]) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(new(T), id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
