package service

import (
	"context"
	"tronopay/internal/entity"

	"github.com/google/uuid"
)

type TaskService struct {
	deps *ServicesDependencies
}

func NewTaskService(deps *ServicesDependencies) *TaskService {
	return &TaskService{deps: deps}
}

func (s *TaskService) Create(ctx context.Context, e entity.Task) error {
	return s.deps.Repos.Task.Create(ctx, e)
}

func (s *TaskService) Delete(ctx context.Context, guid uuid.UUID) error {
	return s.deps.Repos.Task.Delete(ctx, guid)
}

func (s *TaskService) GetTaskById(ctx context.Context, guid uuid.UUID) (entity.Task, error) {
	return s.deps.Repos.Task.GetById(ctx, guid)
}

func (s *TaskService) GetTaskList(ctx context.Context) ([]entity.Task, error) {
	return s.deps.Repos.Task.GetList(ctx)
}
