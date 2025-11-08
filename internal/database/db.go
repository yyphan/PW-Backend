package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"

	"yyphan-pw/backend/internal/models"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error reading .env file")
	}

	dsn := getDSNFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected")
}

func ValidateDBModels() {
	if err := DB.AutoMigrate(models.GetAllDBModels()...); err != nil {
		log.Fatalf("Db model mismatch: %v", err)
	}

	log.Println("Db model validated")
}

func getDSNFromEnv() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || pass == "" || name == "" {
		log.Fatal("Missing DB environment variables in .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		pass,
		name,
		port,
	)

	return dsn
}
