package main

import (
	"log"
	"os"

	_ "Basketball_academy/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Basketball_academy/models"
	"Basketball_academy/routes"
)

// @title Basketball Academy API
// @version 1.0
// @description API для Basketball Academy
// @host localhost:8080
// @BasePath /
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	err = db.AutoMigrate(
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

	routes.SetupRoutes(router, db)

	// Swagger будет доступен по адресу: http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	router.Run("localhost:8080")
}
