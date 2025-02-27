package database

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB подключается к базе данных и возвращает экземпляр *gorm.DB
func ConnectDB() *gorm.DB {
	log.Println("Начало подключения к базе данных...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}
	log.Println("Переменные окружения успешно загружены")

	// Выводим значения переменных окружения для отладки
	log.Printf("Host: %s", os.Getenv("POSTGRESQL_HOST"))
	log.Printf("User: %s", os.Getenv("POSTGRESQL_USER"))
	log.Printf("Database: %s", os.Getenv("POSTGRESQL_DBNAME"))
	log.Printf("Port: %s", os.Getenv("POSTGRESQL_PORT"))

	// Получаем пароль из переменной окружения и кодируем его для URL
	password := "0,zU>9Ko$u:$N\\"
	encodedPassword := url.QueryEscape(password)

	// Используем URL-формат для строки подключения
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=UTC",
		os.Getenv("POSTGRESQL_USER"),
		encodedPassword,
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DBNAME"),
	)

	log.Println("====== СТРОКА ПОДКЛЮЧЕНИЯ ======")
	log.Printf("DSN: %s", dsn)
	log.Println("===============================")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	log.Println("Подключение к БД успешно установлено")
	return db
}
