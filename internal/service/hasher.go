package service

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasherService struct{}

func NewBcryptHasherService() BcryptHasherService {
	return BcryptHasherService{}
}

func (ph BcryptHasherService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ph BcryptHasherService) Compare(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
