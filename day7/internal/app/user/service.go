package user

import (
	"context"

	"github.com/harisfi/alterra-agmc/day7/internal/dto"
	"github.com/harisfi/alterra-agmc/day7/internal/factory"
	"github.com/harisfi/alterra-agmc/day7/internal/model"
	"github.com/harisfi/alterra-agmc/day7/internal/repository"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	CreateUser(c context.Context, payload *dto.CreateUserRequest) (*model.User, error)
	FindUser(c context.Context) (*[]model.User, error)
	FindUserByID(c context.Context, ID uint) (*model.User, error)
	UpdateUser(c context.Context, ID uint, payload *dto.UpdateUserRequest) (*model.User, error)
	DeleteUser(c context.Context, ID uint) error
}

func NewService(f *factory.Factory) Service {
	return &service{UserRepository: f.UserRepository}
}

func (s *service) CreateUser(c context.Context, payload *dto.CreateUserRequest) (*model.User, error) {
	var user = model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := s.UserRepository.CreateUser(c, user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) FindUser(c context.Context) (*[]model.User, error) {
	users, err := s.UserRepository.FindUser(c)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *service) FindUserByID(c context.Context, ID uint) (*model.User, error) {
	user, err := s.UserRepository.FindUserByID(c, ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) UpdateUser(c context.Context, ID uint, payload *dto.UpdateUserRequest) (*model.User, error) {
	user, err := s.UserRepository.FindUserByID(c, ID)

	if err != nil {
		return nil, err
	}

	if payload.Name != nil {
		user.Name = *payload.Name
	}
	if payload.Email != nil {
		user.Email = *payload.Email
	}
	if payload.Password != nil {
		user.Password = *payload.Password
	}

	if err := s.UserRepository.UpdateUser(c, ID, user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) DeleteUser(c context.Context, ID uint) error {
	if _, err := s.UserRepository.FindUserByID(c, ID); err != nil {
		return err
	}
	if err := s.UserRepository.DeleteUser(c, ID); err != nil {
		return err
	}

	return nil
}
