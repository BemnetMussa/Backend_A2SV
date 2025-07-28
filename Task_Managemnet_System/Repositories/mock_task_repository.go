package repositories

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByID(ctx context.Context, id string) (domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(ctx context.Context, id string, task domain.Task) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
