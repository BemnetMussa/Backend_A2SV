package infrastructure

import "github.com/stretchr/testify/mock"


// Mock for JWTService:

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(userID, email, role string) (string, error) {
	args := m.Called(userID, email, role)
	return args.String(0), args.Error(1)
}

// Mock for PasswordService:

type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) ComparePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}
