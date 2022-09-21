package repository

import (
	"context"

	"github.com/harisfi/alterra-agmc/day6/internal/model"
	"gorm.io/gorm"
)

type Book interface {
	CreateBook(c context.Context, book model.Book) error
	FindBook(c context.Context) ([]model.Book, error)
	FindBookByID(c context.Context, ID uint) (model.Book, error)
	UpdateBook(c context.Context, ID uint, book model.Book) error
	DeleteBook(c context.Context, ID uint) error
}

type bookConn struct {
	DB *gorm.DB
}

func NewBook(db *gorm.DB) *bookConn {
	return &bookConn{db}
}

func (conn *bookConn) CreateBook(c context.Context, book model.Book) error {
	return conn.DB.WithContext(c).Model(&model.Book{}).Create(&book).Error
}

func (conn *bookConn) FindBook(c context.Context) ([]model.Book, error) {
	var books []model.Book

	err := conn.DB.WithContext(c).Model(&model.Book{}).Find(&books).Error
	return books, err
}

func (conn *bookConn) FindBookByID(c context.Context, ID uint) (model.Book, error) {
	var book model.Book

	err := conn.DB.WithContext(c).Model(&book).Where("id = ?", ID).First(&book).Error
	return book, err
}

func (conn *bookConn) UpdateBook(c context.Context, ID uint, book model.Book) error {
	err := conn.DB.WithContext(c).Where("id = ?", ID).Model(&model.Book{}).Updates(book).Error
	return err
}

func (conn *bookConn) DeleteBook(c context.Context, ID uint) error {
	err := conn.DB.WithContext(c).Where("id = ?", ID).Delete(&model.Book{}).Error
	return err
}
