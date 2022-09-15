package database

import (
	"github.com/harisfi/alterra-agmc/day3/configs"
	"github.com/harisfi/alterra-agmc/day3/middlewares"
	"github.com/harisfi/alterra-agmc/day3/models"
)

func GetAllUsers() (interface{}, error) {
	var users []models.User

	if e := configs.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUserById(id uint) (interface{}, error) {
	var user models.User
	user.ID = id

	if e := configs.DB.First(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if e := configs.DB.Save(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func UpdateUserById(id uint, user models.User) (interface{}, error) {
	_, e := GetUserById(id)
	if e != nil {
		return nil, e
	}

	user.ID = id
	if e := configs.DB.Updates(&user).Error; e != nil {
		return nil, e
	}

	return GetUserById(id)
}

func DeleteUserById(id uint) error {
	_, e := GetUserById(id)
	if e != nil {
		return e
	}

	if e := configs.DB.Delete(&models.User{}, id).Error; e != nil {
		return e
	}

	return nil
}

func LoginUser(user *models.User) (interface{}, error) {
	if e := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; e != nil {
		return nil, e
	}

	token, e := middlewares.CreateToken(int(user.ID))
	if e != nil {
		return nil, e
	}

	return token, nil
}
