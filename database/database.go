package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ensureEnvFile проверяет, существует ли .env, и если нет – создаёт его,
// используя значения из переменных окружения.
func ensureEnvFile() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		log.Println(".env файл не найден, создаём его...")
		content := fmt.Sprintf(
			"POSTGRESQL_HOST=%s\nPOSTGRESQL_PORT=%s\nPOSTGRESQL_USER=%s\nPOSTGRESQL_PASSWORD=%s\nPOSTGRESQL_DBNAME=%s\n",
			os.Getenv("POSTGRESQL_HOST"),
			os.Getenv("POSTGRESQL_PORT"),
			os.Getenv("POSTGRESQL_USER"),
			os.Getenv("POSTGRESQL_PASSWORD"),
			os.Getenv("POSTGRESQL_DBNAME"),
		)
		if err := ioutil.WriteFile(".env", []byte(content), 0644); err != nil {
			log.Fatal("Ошибка создания .env файла:", err)
		}
		log.Println(".env файл успешно создан")
	}
}

// ConnectDB подключается к базе данных и возвращает экземпляр *gorm.DB
func ConnectDB() *gorm.DB {
	ensureEnvFile()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}
	log.Println("Переменные окружения успешно загружены")

	// Выводим значения переменных окружения для отладки
	log.Printf("Host: %s", os.Getenv("POSTGRESQL_HOST"))
	log.Printf("User: %s", os.Getenv("POSTGRESQL_USER"))
	log.Printf("Database: %s", os.Getenv("POSTGRESQL_DBNAME"))
	log.Printf("Port: %s", os.Getenv("POSTGRESQL_PORT"))

	// Получаем пароль из переменной окружения и кодируем его для URL.
	// Здесь можно использовать жестко заданный пароль для отладки,
	// либо брать его сразу из .env (он уже установлен в переменных окружения):
	password := "0,zU>9Ko$u:$N\\"
	encodedPassword := url.QueryEscape(password)
	log.Printf("Закодированный пароль: %s", encodedPassword)

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
