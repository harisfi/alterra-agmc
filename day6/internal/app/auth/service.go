package auth

import (
	"context"

	"github.com/harisfi/alterra-agmc/day6/internal/dto"
	"github.com/harisfi/alterra-agmc/day6/internal/factory"
	"github.com/harisfi/alterra-agmc/day6/internal/model"
	"github.com/harisfi/alterra-agmc/day6/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	Repository repository.User
}

type Service interface {
	Login(c context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(c context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{Repository: f.UserRepository}
}

func (s *service) Login(c context.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	user, err := s.Repository.FindUserByEmail(c, payload.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, err
	}

	result = &dto.AuthLoginResponse{
		Token: token,
		User:  user,
	}

	return result, nil
}

func (s *service) Register(c context.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse
	var user model.User = payload.User

	if err := s.Repository.CreateUser(c, user); err != nil {
		return nil, err
	}

	result = &dto.AuthRegisterResponse{
		User: user,
	}

	return result, nil
}
