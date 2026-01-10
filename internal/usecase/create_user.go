package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
)

type UserRepository interface {
	Save(*entity.User) error
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

func (uc *CreateUserUseCase) Execute(username, password string) error {
	if username == "" || len(username) < 4 || len(username) > 50 {
		return fmt.Errorf("Логин должен содержать от 4 до 50 символов")
	}

	if password == "" || len(password) < 8 || len(password) > 20 {
		return fmt.Errorf("Пароль должен содержать от 4 до 20 символов")
	}

	hashPassword, err := uc.ph.Hash(password)
	if err != nil {
		return fmt.Errorf("Ошибка при хешировании пароля: %s", err.Error())
	}

	u := entity.User{Username: username, Password: hashPassword}

	if err = uc.r.Save(&u); err != nil {
		// todo проверка на unique constraint username
		return fmt.Errorf("Ошибка при создании пользователя: %s", err.Error())
	}

	return nil
}
