package database

import (
	"github.com/harisfi/alterra-agmc/day4/learn-echo-test/configs"
	"github.com/harisfi/alterra-agmc/day4/learn-echo-test/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := configs.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}
