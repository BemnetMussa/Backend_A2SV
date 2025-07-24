package usecases

import (
	"context"
	"errors"
	"time"
	
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/repositories"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo       repositories.UserRepository
	JWTService *infrastructure.JWTService
}


func NewUserUsecase(repo repositories.UserRepository, jwtService *infrastructure.JWTService) *UserUsecase {
	return &UserUsecase{Repo: repo, JWTService: jwtService}
}


func (uc *UserUsecase) RegisterUser(name, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := uc.Repo.FindByEmail(ctx, email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	count, err := uc.Repo.CountUsers(ctx)
	if err != nil {
		return errors.New("failed to count users")
	}

	role := "user"
	if count == 0 {
		role = "admin"
	}

	newUser := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = uc.Repo.CreateUser(ctx, newUser)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}


func (uc *UserUsecase) LoginUser(email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := uc.Repo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Use JWTService to generate token
	token, err := uc.JWTService.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}



func (uc *UserUsecase) PromoteUserByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return uc.Repo.PromoteUserByEmail(ctx, email)
}
