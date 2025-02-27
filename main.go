package main

import (
	"Basketball_academy/database"
	_ "Basketball_academy/docs"
	"Basketball_academy/models"
	"Basketball_academy/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("%s %s %d %s", c.Request.Method, c.Request.URL, c.Writer.Status(), time.Since(start))
	}
}

func main() {
	log.Println("Запуск приложения...")

	// Подключаемся к базе данных
	db := database.ConnectDB()

	err := db.AutoMigrate(
		&models.User{},
		&models.Trainer{},
		&models.News{},
		&models.Promotion{},
		&models.TrainingSession{},
		&models.TrainingRegistration{},
		&models.Chat{},
		&models.ChatMessage{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
		&models.Notification{},
		&models.PromoCode{},
		&models.SupportTicket{},
		&models.AcademyStaff{},
		&models.TrainingGroup{},
		&models.MyTraining{},
		&models.Cart{},
		&models.CartItem{},
		&models.Club{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	// Инициализируем Gin
	router := gin.Default()

	// Применяем middleware для логирования запросов
	router.Use(LoggerMiddleware())

	routes.SetupRoutes(router, db)

	// Swagger будет доступен по адресу: http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	router.Run("localhost:8080")
}
