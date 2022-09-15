package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/harisfi/alterra-agmc/day4/learn-echo-test/models"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_UNAME": "root",
		"DB_PASS":  "",
		"DB_HOST":  "127.0.0.1",
		"DB_PORT":  "3306",
		"DB_NAME":  "test",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_UNAME"],
		config["DB_PASS"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
