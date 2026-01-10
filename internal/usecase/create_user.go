package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
)

type UserRepository interface {
	CreateUser(*entity.User) error
	FindByUsername(string) (entity.User, error)
}

type PasswordHasher interface {
	Hash(string) (string, error)
	Compare(string, string) bool
}

type CreateUserUseCase struct {
	r  UserRepository
	ph PasswordHasher
}

func NewCreateUserUseCase(repo UserRepository, hasher PasswordHasher) CreateUserUseCase {
	return CreateUserUseCase{
		r:  repo,
		ph: hasher,
	}
}

func (uc CreateUserUseCase) Execute(username, password string) error {
	hashPassword, err := uc.ph.Hash(password)
	if err != nil {
	}

	fmt.Println(hashPassword)

	return nil
}
