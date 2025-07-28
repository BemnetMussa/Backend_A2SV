package usecases

import (
	"context"
	"time"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/repositories"
)

type TaskUsecase struct {
	Repo repositories.TaskRepository
}

func NewTaskUsecase(repo repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{Repo: repo}
}

func (uc *TaskUsecase) CreateTask(task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.Repo.Create(ctx, task)
}

func (uc *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.Repo.GetAll(ctx)
}

func (uc *TaskUsecase) GetTaskByID(id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.Repo.GetByID(ctx, id)
}

func (uc *TaskUsecase) UpdateTask(id string, task domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.Repo.Update(ctx, id, task)
}

func (uc *TaskUsecase) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.Repo.Delete(ctx, id)
}
