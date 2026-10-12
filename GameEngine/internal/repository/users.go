
package repository

import (
	"context"
	"errors"

	"github.com/iDIOtu/PivchikiDND/internal/database/sqlc"
	"github.com/iDIOtu/PivchikiDND/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context) (*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	List(ctx context.Context, limit, offset int32) ([]domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) Create(ctx context.Context) (*domain.User, error) {
	row, err := r.q.CreateUser(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        row.ID,
		CreatedAt:          row.CreatedAt.Time,
	}, nil
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.User, error) {
	row, err := r.q.GetUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        row.ID,
		CreatedAt:          row.CreatedAt.Time,
	}, nil
}

func (r *userRepository) List(
	ctx context.Context,
	limit, offset int32,
) ([]domain.User, error) {
	if limit <= 0 || offset < 0 {
		return nil, ErrInvalidInput
	}

	rows, err := r.q.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(rows))
	for _, row := range rows {
		users = append(users, domain.User{
			ID:        row.ID,
		CreatedAt:          row.CreatedAt.Time,
		})
	}

	return users, nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	rows, err := r.q.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
