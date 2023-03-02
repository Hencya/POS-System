package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"POS-System/helpers"
	"POS-System/repository/database/transactionRepo"
	"POS-System/repository/database/userRepo"
)

func SetupDatabaseConnection() *gorm.DB {
	_ = godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err.Error())
	}
	dbMigrate(db)
	dbSeed(db)

	return db
}

func dbMigrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(
		&userRepo.User{},
		&transactionRepo.Transaction{},
	)
}

func dbSeed(db *gorm.DB) {
	password, _ := helpers.HashPassword("password")
	db.Create(&userRepo.User{ID: 1, Username: "username", Password: password})
}
