package usecase

import (
	"errors"
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/domain"
)

type UserSaver interface {
	Save(*domain.User) error
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
		return domain.ErrInvalidUsername
	}

	if password == "" || len(password) < 8 || len(password) > 20 {
		return domain.ErrInvalidPassword
	}

	hashedPassword, err := uc.ph.Hash(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	u := domain.User{
		Username: username,
		Password: hashedPassword,
	}

	if err = uc.us.Save(&u); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return domain.ErrUserAlreadyExists
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
