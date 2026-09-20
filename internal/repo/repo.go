package repo

import (
	"context"

	"tronopay/internal/entity"
	"tronopay/internal/repo/pgdb"
	"tronopay/pkg/postgres"

	"github.com/google/uuid"
)

type Task interface {
	Create(ctx context.Context, entity entity.Task) error
	Delete(ctx context.Context, guid uuid.UUID) error
	GetById(ctx context.Context, guid uuid.UUID) (entity.Task, error)
	GetList(ctx context.Context) ([]entity.Task, error)
}

type (
	Repositories struct {
		Task
		// ...
	}
)

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Task: pgdb.NewRepoTask(pg),
		// ...
	}
}
