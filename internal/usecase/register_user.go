package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
)

type UserSaver interface {
	Save(*entity.User) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type RegisterUser struct {
	us UserSaver
	ph PasswordHasher
}

func NewRegisterUser(us UserSaver, hasher PasswordHasher) *RegisterUser {
	return &RegisterUser{
		us: us,
		ph: hasher,
	}
}

func (uc *RegisterUser) Execute(username, password string) error {
	if username == "" || len(username) < 4 || len(username) > 50 {
		return fmt.Errorf("Логин должен содержать от 4 до 50 символов")
	}

	if password == "" || len(password) < 8 || len(password) > 20 {
		return fmt.Errorf("Пароль должен содержать от 8 до 20 символов")
	}

	hashPassword, err := uc.ph.Hash(password)
	if err != nil {
		return fmt.Errorf("Ошибка при хешировании пароля: %s", err.Error())
	}

	u := entity.User{Username: username, Password: hashPassword}

	if err = uc.us.Save(&u); err != nil {
		// todo проверка на unique constraint username
		return fmt.Errorf("Ошибка при создании пользователя: %s", err.Error())
	}

	return nil
}
