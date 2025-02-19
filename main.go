package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Импортируем сгенерированную документацию Swagger
	_ "Basketball_academy/docs"

	// Импортируем Swagger UI middleware
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"Basketball_academy/models"
	"Basketball_academy/routes"
)

// @title Basketball Academy API
// @version 1.0
// @description API для Basketball Academy
// @host localhost:8080
// @BasePath /
func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Формируем строку подключения (DSN)
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=UTC"

	// Подключаемся к БД
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	// Автоматическая миграция для всех моделей
	err = db.AutoMigrate(
		&models.User{},
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
		&models.StaticPage{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	// Инициализируем Gin
	router := gin.Default()

	// Регистрируем маршруты вашего API
	routes.SetupRoutes(router, db)

	// Регистрируем маршрут для Swagger UI
	// Swagger будет доступен по адресу: http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	router.Run(":8080")
}
