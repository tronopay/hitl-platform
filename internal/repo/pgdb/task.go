package pgdb

import (
	"context"
	"errors"
	"fmt"

	"tronopay/internal/entity"
	"tronopay/internal/repo/repoerrors"
	"tronopay/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RepoTask struct {
	*postgres.Postgres
}

func NewRepoTask(pg *postgres.Postgres) *RepoTask {
	return &RepoTask{pg}
}

func (r *RepoTask) Create(ctx context.Context, entity entity.Task) error {
	sql, args, _ := r.Builder.
		Insert("tasks").
		Columns("id", "status", "type", "description").
		Values(entity.Id, entity.Status, entity.Type, entity.Description).
		ToSql()

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return repoerrors.ErrNotInserted
	}

	return nil
}

func (r *RepoTask) Delete(ctx context.Context, guid uuid.UUID) error {
	sql, args, _ := r.Builder.
		Delete("tasks").
		Where("id = $1", guid).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return repoerrors.ErrNotDeleted
	}

	return nil
}

func (r *RepoTask) GetById(ctx context.Context, guid uuid.UUID) (entity.Task, error) {
	sql, args, _ := r.Builder.
		Select("id, created_at, updated_at, status, type, description").
		From("tasks").
		Where("id = $1", guid).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var record entity.Task
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&record.Id,
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.Status,
		&record.Type,
		&record.Description,
	)

	if err == nil {
		return record, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Task{}, repoerrors.ErrNotFound
	}

	return entity.Task{}, fmt.Errorf("RepoTask.GetById - r.Pool.QueryRow: %v", err)
}

func (r *RepoTask) GetList(ctx context.Context) ([]entity.Task, error) {
	sql, args, _ := r.Builder.
		Select("id, created_at, updated_at, status, type, description").
		From("tasks").
		OrderBy("id ASC").
		ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("RepoTask.GetList - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var records []entity.Task
	for rows.Next() {
		var record entity.Task
		err = rows.Scan(
			&record.Id,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.Status,
			&record.Type,
			&record.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("RepoTask.GetList - rows.Scan: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}
