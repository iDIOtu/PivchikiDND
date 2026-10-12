
package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/iDIOtu/PivchikiDND/internal/database/sqlc"
	"github.com/iDIOtu/PivchikiDND/internal/domain"
	"github.com/jackc/pgx/v5"
)

type GameRepository interface {
	Create(ctx context.Context, game *domain.Game) (*domain.Game, error)
	GetByID(ctx context.Context, id int64) (*domain.Game, error)
	ListByUser(ctx context.Context, userID int64, limit, offset int32) ([]domain.Game, error)
	Update(ctx context.Context, game *domain.Game) (*domain.Game, error)
	Delete(ctx context.Context, id int64) error
}

type gameRepository struct {
	q *sqlc.Queries
}

func NewGameRepository(q *sqlc.Queries) GameRepository {
	return &gameRepository{q: q}
}

func (r *gameRepository) Create(
	ctx context.Context,
	game *domain.Game,
) (*domain.Game, error) {
	if game == nil || game.UserID <= 0 || game.Name == "" {
		return nil, ErrInvalidInput
	}
	if !json.Valid(game.Ruleset) {
		return nil, ErrInvalidJSON
	}

	row, err := r.q.CreateGame(ctx, sqlc.CreateGameParams{
		UserID:             game.UserID,
		Name:               game.Name,
		SettingDescription: game.SettingDescription,
		Ruleset:            game.Ruleset,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Game{
		ID:                 row.ID,
		UserID:             row.UserID,
		Name:               row.Name,
		SettingDescription: row.SettingDescription,
		Ruleset:            json.RawMessage(row.Ruleset),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *gameRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Game, error) {
	row, err := r.q.GetGame(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.Game{
		ID:                 row.ID,
		UserID:             row.UserID,
		Name:               row.Name,
		SettingDescription: row.SettingDescription,
		Ruleset:            json.RawMessage(row.Ruleset),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *gameRepository) ListByUser(
	ctx context.Context,
	userID int64,
	limit, offset int32,
) ([]domain.Game, error) {
	if userID <= 0 || limit <= 0 || offset < 0 {
		return nil, ErrInvalidInput
	}

	rows, err := r.q.ListGamesByUser(ctx, sqlc.ListGamesByUserParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	games := make([]domain.Game, 0, len(rows))
	for _, row := range rows {
		games = append(games, domain.Game{
			ID:                 row.ID,
			UserID:             row.UserID,
			Name:               row.Name,
			SettingDescription: row.SettingDescription,
			Ruleset:            json.RawMessage(row.Ruleset),
			CreatedAt:          row.CreatedAt.Time,
			UpdatedAt:          row.UpdatedAt.Time,
		})
	}

	return games, nil
}

func (r *gameRepository) Update(
	ctx context.Context,
	game *domain.Game,
) (*domain.Game, error) {
	if game == nil || game.ID <= 0 || game.Name == "" {
		return nil, ErrInvalidInput
	}
	if !json.Valid(game.Ruleset) {
		return nil, ErrInvalidJSON
	}

	row, err := r.q.UpdateGame(ctx, sqlc.UpdateGameParams{
		ID:                 game.ID,
		Name:               game.Name,
		SettingDescription: game.SettingDescription,
		Ruleset:            game.Ruleset,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.Game{
		ID:                 row.ID,
		UserID:             row.UserID,
		Name:               row.Name,
		SettingDescription: row.SettingDescription,
		Ruleset:            json.RawMessage(row.Ruleset),
		CreatedAt:          row.CreatedAt.Time,
		UpdatedAt:          row.UpdatedAt.Time,
	}, nil
}

func (r *gameRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	rows, err := r.q.DeleteGame(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
