package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	mysql := mysqlConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
