package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/domain"
)

type UserFinder interface {
	FindByUsername(username string) (*domain.User, error)
}

type PasswordComparer interface {
	Compare(password, hashedPassword string) bool
}

type AuthorizeUser struct {
	uf UserFinder
	pc PasswordComparer
}

func NewAuthorizeUser(uf UserFinder, pc PasswordComparer) *AuthorizeUser {
	return &AuthorizeUser{
		uf: uf,
		pc: pc,
	}
}

func (uc *AuthorizeUser) Execute(username, password string) error {
	u, err := uc.uf.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("Failed find user by username")
	}

	if u == nil {
		return domain.ErrUserNotFound
	}

	if !uc.pc.Compare(password, u.Password) {
		return domain.ErrInvalidPassword
	}

	return nil
}
