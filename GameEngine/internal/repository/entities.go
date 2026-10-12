
package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/iDIOtu/PivchikiDND/internal/database/sqlc"
	"github.com/iDIOtu/PivchikiDND/internal/domain"
	"github.com/jackc/pgx/v5"
)

type EntityRepository interface {
	Create(ctx context.Context, entity *domain.Entity) (*domain.Entity, error)
	GetByID(ctx context.Context, id int64) (*domain.Entity, error)
	ListByGame(ctx context.Context, gameID int64, limit, offset int32) ([]domain.Entity, error)
	Update(ctx context.Context, entity *domain.Entity) (*domain.Entity, error)
	Delete(ctx context.Context, id int64) error
}

type entityRepository struct {
	q *sqlc.Queries
}

func NewEntityRepository(q *sqlc.Queries) EntityRepository {
	return &entityRepository{q: q}
}

func (r *entityRepository) Create(
	ctx context.Context,
	entity *domain.Entity,
) (*domain.Entity, error) {
	if entity == nil || entity.GameID <= 0 ||
		entity.Type == "" || entity.Name == "" {
		return nil, ErrInvalidInput
	}
	if !json.Valid(entity.Data) {
		return nil, ErrInvalidJSON
	}

	row, err := r.q.CreateEntity(ctx, sqlc.CreateEntityParams{
		GameID: entity.GameID,
		Type:   entity.Type,
		Name:   entity.Name,
		Data:   entity.Data,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Entity{
		ID:        row.ID,
		GameID:    row.GameID,
		Type:      row.Type,
		Name:      row.Name,
		Data:      json.RawMessage(row.Data),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *entityRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Entity, error) {
	row, err := r.q.GetEntity(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.Entity{
		ID:        row.ID,
		GameID:    row.GameID,
		Type:      row.Type,
		Name:      row.Name,
		Data:      json.RawMessage(row.Data),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *entityRepository) ListByGame(
	ctx context.Context,
	gameID int64,
	limit, offset int32,
) ([]domain.Entity, error) {
	if gameID <= 0 || limit <= 0 || offset < 0 {
		return nil, ErrInvalidInput
	}

	rows, err := r.q.ListEntitiesByGame(ctx, sqlc.ListEntitiesByGameParams{
		GameID: gameID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	entities := make([]domain.Entity, 0, len(rows))
	for _, row := range rows {
		entities = append(entities, domain.Entity{
			ID:        row.ID,
			GameID:    row.GameID,
			Type:      row.Type,
			Name:      row.Name,
			Data:      json.RawMessage(row.Data),
			CreatedAt:          row.CreatedAt.Time,
			UpdatedAt:          row.UpdatedAt.Time,
		})
	}

	return entities, nil
}

func (r *entityRepository) Update(
	ctx context.Context,
	entity *domain.Entity,
) (*domain.Entity, error) {
	if entity == nil || entity.ID <= 0 ||
		entity.Type == "" || entity.Name == "" {
		return nil, ErrInvalidInput
	}
	if !json.Valid(entity.Data) {
		return nil, ErrInvalidJSON
	}

	row, err := r.q.UpdateEntity(ctx, sqlc.UpdateEntityParams{
		ID:   entity.ID,
		Type: entity.Type,
		Name: entity.Name,
		Data: entity.Data,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.Entity{
		ID:        row.ID,
		GameID:    row.GameID,
		Type:      row.Type,
		Name:      row.Name,
		Data:      json.RawMessage(row.Data),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *entityRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	rows, err := r.q.DeleteEntity(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
