package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type PasswordServiceInterface interface {
    HashPassword(password string) (string, error)
    ComparePassword(hashedPassword, password string) error
}

type BcryptPasswordService struct{}

func NewPasswordService() *BcryptPasswordService {
	return &BcryptPasswordService{}
}

// HashPassword generates a bcrypt hash from the plaintext password
func (s *BcryptPasswordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ComparePassword checks if the given password matches the hash
func (s *BcryptPasswordService) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
