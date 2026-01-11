package usecase

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
)

type UserFinder interface {
	FindByUsername(username string) (*entity.User, error)
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
		return fmt.Errorf("User not found")
	}

	if !uc.pc.Compare(password, u.Password) {
		return fmt.Errorf("Failed to authorize")
	}

	return nil
}
