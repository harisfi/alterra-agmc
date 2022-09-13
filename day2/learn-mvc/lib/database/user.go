package database

import (
	"github.com/harisfi/alterra-agmc/day2/learn-mvc/configs"
	"github.com/harisfi/alterra-agmc/day2/learn-mvc/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := configs.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
