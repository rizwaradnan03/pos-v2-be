package db

import (
	"fmt"
	"log"
	"os"
	"pos-v2-be/internal/pkg"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		user,
		name,
		port,
	)

	if password != "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			host,
			user,
			password,
			name,
			port,
		)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	for _, md := range pkg.ModelsToMigrate() {
		err = database.AutoMigrate(md)

		if err != nil {
			log.Fatal("Gagal migrasi:", err)
		} else {
			fmt.Println("Berhasil Migrasi")
		}
	}

	DB = database
	fmt.Println("Database connected successfully")
}
