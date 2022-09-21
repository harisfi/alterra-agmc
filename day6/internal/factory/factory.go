package factory

import (
	"github.com/harisfi/alterra-agmc/day6/database"
	"github.com/harisfi/alterra-agmc/day6/internal/repository"
)

type Factory struct {
	BookRepository repository.Book
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		BookRepository: repository.NewBook(db),
		UserRepository: repository.NewUser(db),
	}
}
