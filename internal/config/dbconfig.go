package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
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
	id := "dale"
	password := "0000"
	host := "localhost"
	port := "20003"

	connectionUrl := id + ":" + password + "@tcp(" + host + ":" + port + ")/coin?charset=utf8mb4&parseTime=True&loc=Local"

	return &connectionUrl
}
