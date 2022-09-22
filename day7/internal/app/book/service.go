package book

import (
	"context"

	"github.com/harisfi/alterra-agmc/day7/internal/dto"
	"github.com/harisfi/alterra-agmc/day7/internal/factory"
	"github.com/harisfi/alterra-agmc/day7/internal/model"
	"github.com/harisfi/alterra-agmc/day7/internal/repository"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	CreateBook(c context.Context, payload *dto.CreateBookRequest) (*model.Book, error)
	FindBook(c context.Context) (*[]model.Book, error)
	FindBookByID(c context.Context, ID uint) (*model.Book, error)
	UpdateBook(c context.Context, ID uint, payload *dto.UpdateBookRequest) (*model.Book, error)
	DeleteBook(c context.Context, ID uint) error
}

func NewService(f *factory.Factory) Service {
	return &service{BookRepository: f.BookRepository}
}

func (s *service) CreateBook(c context.Context, payload *dto.CreateBookRequest) (*model.Book, error) {
	var book = model.Book{
		Title:     payload.Title,
		Author:    payload.Author,
		Publisher: payload.Publisher,
	}

	if err := s.BookRepository.CreateBook(c, book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *service) FindBook(c context.Context) (*[]model.Book, error) {
	books, err := s.BookRepository.FindBook(c)

	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (s *service) FindBookByID(c context.Context, ID uint) (*model.Book, error) {
	book, err := s.BookRepository.FindBookByID(c, ID)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *service) UpdateBook(c context.Context, ID uint, payload *dto.UpdateBookRequest) (*model.Book, error) {
	book, err := s.BookRepository.FindBookByID(c, ID)

	if err != nil {
		return nil, err
	}

	if payload.Title != nil {
		book.Title = *payload.Title
	}
	if payload.Author != nil {
		book.Author = *payload.Author
	}
	if payload.Publisher != nil {
		book.Publisher = *payload.Publisher
	}

	if err := s.BookRepository.UpdateBook(c, ID, book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *service) DeleteBook(c context.Context, ID uint) error {
	if _, err := s.BookRepository.FindBookByID(c, ID); err != nil {
		return err
	}
	if err := s.BookRepository.DeleteBook(c, ID); err != nil {
		return err
	}

	return nil
}
