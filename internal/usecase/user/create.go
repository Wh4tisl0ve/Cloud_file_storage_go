package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
)

type UserRepository interface {
	CreateUser(u *entity.User) error
	FindByUsername(userName string) (entity.User, error)
}

type CreateUserUseCase struct {
	r UserRepository
}

func NewCreateUserUseCase(repo UserRepository) CreateUserUseCase{
	return CreateUserUseCase{
		r: repo,
	}
}

func (uc CreateUserUseCase) Execute(username, password string) error {
	fmt.Println("user created!")
	return nil
}
