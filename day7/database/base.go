package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConfig struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
