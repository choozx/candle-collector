package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := getDBUrl()
	DB, err = gorm.Open(mysql.Open(*dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %w", err)
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("failed to close database connection: %v", err)
		return
	}
	sqlDB.Close()
}

func getDBUrl() *string {
	id := os.Getenv("DB_ID")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connectionUrl := id + ":" + password + "@tcp(" + host + ":" + port + ")/coin?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println(connectionUrl)

	return &connectionUrl
}
