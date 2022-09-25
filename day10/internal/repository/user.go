package repository

import (
	"context"

	"github.com/harisfi/alterra-agmc/day10/internal/model"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(c context.Context, user model.User) error
	FindUser(c context.Context) ([]model.User, error)
	FindUserByID(c context.Context, ID uint) (model.User, error)
	FindUserByEmail(c context.Context, email string) (model.User, error)
	UpdateUser(c context.Context, ID uint, user model.User) error
	DeleteUser(c context.Context, ID uint) error
}

type userConn struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *userConn {
	return &userConn{db}
}

func (conn *userConn) CreateUser(c context.Context, user model.User) error {
	return conn.DB.WithContext(c).Model(&model.User{}).Create(&user).Error
}

func (conn *userConn) FindUser(c context.Context) ([]model.User, error) {
	var users []model.User

	err := conn.DB.WithContext(c).Model(&model.User{}).Find(&users).Error
	return users, err
}

func (conn *userConn) FindUserByID(c context.Context, ID uint) (model.User, error) {
	var user model.User

	err := conn.DB.WithContext(c).Model(&user).Where("id = ?", ID).First(&user).Error
	return user, err
}

func (conn *userConn) FindUserByEmail(c context.Context, email string) (model.User, error) {
	var user model.User

	err := conn.DB.WithContext(c).Model(&user).Where("email = ?", email).First(&user).Error
	return user, err
}

func (conn *userConn) UpdateUser(c context.Context, ID uint, user model.User) error {
	err := conn.DB.WithContext(c).Where("id = ?", ID).Model(&model.User{}).Updates(user).Error
	return err
}

func (conn *userConn) DeleteUser(c context.Context, ID uint) error {
	err := conn.DB.WithContext(c).Where("id = ?", ID).Delete(&model.User{}).Error
	return err
}
