package repositories

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*domain.User), args.Error(1)
}


func (m *MockUserRepository) CountUsers(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) PromoteUserByEmail(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}
