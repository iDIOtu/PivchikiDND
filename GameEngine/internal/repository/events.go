
package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/iDIOtu/PivchikiDND/internal/database/sqlc"
	"github.com/iDIOtu/PivchikiDND/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type EventRepository interface {
	Create(ctx context.Context, event *domain.Event) (*domain.Event, error)
	GetByID(ctx context.Context, id int64) (*domain.Event, error)
	ListByGame(ctx context.Context, gameID int64, limit, offset int32) ([]domain.Event, error)
	Delete(ctx context.Context, id int64) error
}

type eventRepository struct {
	q *sqlc.Queries
}

func NewEventRepository(q *sqlc.Queries) EventRepository {
	return &eventRepository{q: q}
}

func int64ToPgtype(value *int64) pgtype.Int8 {
	if value == nil {
		return pgtype.Int8{}
	}

	return pgtype.Int8{
		Int64: *value,
		Valid: true,
	}
}

func pgtypeToInt64(value pgtype.Int8) *int64 {
	if !value.Valid {
		return nil
	}

	result := value.Int64
	return &result
}

func (r *eventRepository) Create(
	ctx context.Context,
	event *domain.Event,
) (*domain.Event, error) {
	if event == nil || event.GameID <= 0 || event.Type == "" {
		return nil, ErrInvalidInput
	}
	if !json.Valid(event.Payload) {
		return nil, ErrInvalidJSON
	}

	row, err := r.q.CreateEvent(ctx, sqlc.CreateEventParams{
		GameID:   event.GameID,
		Type:     event.Type,
		ActorID:  int64ToPgtype(event.ActorID),
		TargetID: int64ToPgtype(event.TargetID),
		Payload:  event.Payload,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Event{
		ID:        row.ID,
		GameID:    row.GameID,
		Type:      row.Type,
		ActorID:   pgtypeToInt64(row.ActorID),
		TargetID:  pgtypeToInt64(row.TargetID),
		Payload:   json.RawMessage(row.Payload),
		CreatedAt:          row.CreatedAt.Time,
	}, nil
}

func (r *eventRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Event, error) {
	row, err := r.q.GetEvent(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &domain.Event{
		ID:        row.ID,
		GameID:    row.GameID,
		Type:      row.Type,
		ActorID:   pgtypeToInt64(row.ActorID),
		TargetID:  pgtypeToInt64(row.TargetID),
		Payload:   json.RawMessage(row.Payload),
		CreatedAt:          row.CreatedAt.Time,
	}, nil
}

func (r *eventRepository) ListByGame(
	ctx context.Context,
	gameID int64,
	limit, offset int32,
) ([]domain.Event, error) {
	if gameID <= 0 || limit <= 0 || offset < 0 {
		return nil, ErrInvalidInput
	}

	rows, err := r.q.ListEventsByGame(ctx, sqlc.ListEventsByGameParams{
		GameID: gameID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	events := make([]domain.Event, 0, len(rows))
	for _, row := range rows {
		events = append(events, domain.Event{
			ID:        row.ID,
			GameID:    row.GameID,
			Type:      row.Type,
			ActorID:   pgtypeToInt64(row.ActorID),
			TargetID:  pgtypeToInt64(row.TargetID),
			Payload:   json.RawMessage(row.Payload),
			CreatedAt:          row.CreatedAt.Time,
		})
	}

	return events, nil
}

func (r *eventRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	rows, err := r.q.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
