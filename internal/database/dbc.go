package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDbClient() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass	:= os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			PrepareStmt: true,
		},
	)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	return db, err
}