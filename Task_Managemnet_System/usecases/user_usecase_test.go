package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/infrastructure"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/repositories"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/usecases"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo     *repositories.MockUserRepository
	mockJWT      *infrastructure.MockJWTService
	mockPassword *infrastructure.MockPasswordService
	uc           *usecases.UserUsecase
}

func (s *UserUsecaseTestSuite) SetupTest() {
	s.mockRepo = new(repositories.MockUserRepository)
	s.mockJWT = new(infrastructure.MockJWTService)
	s.mockPassword = new(infrastructure.MockPasswordService)
	s.uc = usecases.NewUserUsecase(s.mockRepo, s.mockJWT, s.mockPassword)
}

func (s *UserUsecaseTestSuite) TestRegisterUser_Success() {
	email, name, password, hashed := "test@example.com", "Test User", "password", "hashedpass"

	s.mockRepo.On("FindByEmail", mock.Anything, email).Return((*domain.User)(nil), errors.New("not found"))

	s.mockPassword.On("HashPassword", password).Return(hashed, nil)
	s.mockRepo.On("CountUsers", mock.Anything).Return(int64(0), nil)
	s.mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	err := s.uc.RegisterUser(name, email, password)
	s.NoError(err)
}

func (s *UserUsecaseTestSuite) TestRegisterUser_UserAlreadyExists() {
	email := "exists@example.com"
	s.mockRepo.On("FindByEmail", mock.Anything, email).Return(&domain.User{}, nil)

	err := s.uc.RegisterUser("Test User", email, "password")
	s.EqualError(err, "user already exists")
}

func (s *UserUsecaseTestSuite) TestRegisterUser_HashPasswordError() {
	email := "test@example.com"
	password := "password"

	s.mockRepo.On("FindByEmail", mock.Anything, email).Return((*domain.User)(nil), errors.New("not found"))

	s.mockPassword.On("HashPassword", password).Return("", errors.New("hash error"))

	err := s.uc.RegisterUser("Test User", email, password)
	s.EqualError(err, "failed to hash password")
}

func (s *UserUsecaseTestSuite) TestRegisterUser_CountUsersError() {
	email, password := "test@example.com", "password"
	s.mockRepo.On("FindByEmail", mock.Anything, email).Return((*domain.User)(nil), errors.New("not found"))

	s.mockPassword.On("HashPassword", password).Return("hashedpass", nil)
	s.mockRepo.On("CountUsers", mock.Anything).Return(int64(0), errors.New("count error"))

	err := s.uc.RegisterUser("Test User", email, password)
	s.EqualError(err, "failed to count users")
}

func (s *UserUsecaseTestSuite) TestRegisterUser_CreateUserError() {
	email, password := "test@example.com", "password"
	s.mockRepo.On("FindByEmail", mock.Anything, email).Return((*domain.User)(nil), errors.New("not found"))

	s.mockPassword.On("HashPassword", password).Return("hashedpass", nil)
	s.mockRepo.On("CountUsers", mock.Anything).Return(int64(1), nil)
	s.mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(errors.New("create error"))

	err := s.uc.RegisterUser("Test User", email, password)
	s.EqualError(err, "failed to create user")
}

func (s *UserUsecaseTestSuite) TestLoginUser_Success() {
	email := "test@example.com"
	password := "password"
	hashed := "hashedpass"
	user := &domain.User{
		ID:       [12]byte{1, 2, 3},
		Email:    email,
		Password: hashed,
		Role:     "user",
	}

	s.mockRepo.On("FindByEmail", mock.Anything, email).Return(user, nil)
	s.mockPassword.On("ComparePassword", hashed, password).Return(nil)
	s.mockJWT.On("GenerateToken", mock.Anything, email, "user").Return("mock-token", nil)

	token, err := s.uc.LoginUser(email, password)
	s.NoError(err)
	s.Equal("mock-token", token)
}

func (s *UserUsecaseTestSuite) TestLoginUser_UserNotFound() {
	email := "notfound@example.com"
	s.mockRepo.On("FindByEmail", mock.Anything, email).Return((*domain.User)(nil), errors.New("not found"))


	token, err := s.uc.LoginUser(email, "password")
	s.Error(err)
	s.Empty(token)
	s.EqualError(err, "invalid email or password")
}

func (s *UserUsecaseTestSuite) TestLoginUser_PasswordCompareFail() {
	email := "test@example.com"
	password := "wrongpass"
	user := &domain.User{Email: email, Password: "hashedpass", Role: "user"}

	s.mockRepo.On("FindByEmail", mock.Anything, email).Return(user, nil)
	s.mockPassword.On("ComparePassword", "hashedpass", password).Return(errors.New("compare fail"))

	token, err := s.uc.LoginUser(email, password)
	s.Error(err)
	s.Empty(token)
	s.EqualError(err, "invalid email or password")
}

func (s *UserUsecaseTestSuite) TestLoginUser_GenerateTokenFail() {
	email := "test@example.com"
	password := "password"
	user := &domain.User{Email: email, Password: "hashedpass", Role: "user"}

	s.mockRepo.On("FindByEmail", mock.Anything, email).Return(user, nil)
	s.mockPassword.On("ComparePassword", "hashedpass", password).Return(nil)
	s.mockJWT.On("GenerateToken", mock.Anything, email, "user").Return("", errors.New("token error"))

	token, err := s.uc.LoginUser(email, password)
	s.Error(err)
	s.Empty(token)
	s.EqualError(err, "failed to generate token")
}

func (s *UserUsecaseTestSuite) TestPromoteUserByEmail_Success() {
	email := "promote@example.com"
	s.mockRepo.On("PromoteUserByEmail", mock.Anything, email).Return(nil)

	err := s.uc.PromoteUserByEmail(email)
	s.NoError(err)
}

func (s *UserUsecaseTestSuite) TestPromoteUserByEmail_Error() {
	email := "promote@example.com"
	s.mockRepo.On("PromoteUserByEmail", mock.Anything, email).Return(errors.New("promote error"))

	err := s.uc.PromoteUserByEmail(email)
	s.Error(err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
