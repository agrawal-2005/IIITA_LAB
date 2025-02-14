package config

import (
	"fmt"
	"log"
	"os"
	"net/url"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/agrawal-2005/iiita_lab/pkg/models"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Panic("❌ Error loading .env file:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		log.Panic("❌ Database credentials missing! Check .env file.")
	}

	escapedPass := url.QueryEscape(dbPass)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		dbUser, escapedPass, dbHost, dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Failed to connect to database:", err)
	}

	models.AutoMigrate(DB)

	fmt.Println("✅ Database connected & migrated successfully!")
}

func GetDB() *gorm.DB {
	return DB
}