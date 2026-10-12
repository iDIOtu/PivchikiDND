package delivery

import (
	dbsqlc "github.com/iDIOtu/PivchikiDND/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type Handler struct {
	db      *pgxpool.Pool
	queries *dbsqlc.Queries
}

func NewHandler(
	db *pgxpool.Pool,
	queries *dbsqlc.Queries,
) *Handler {
	return &Handler{
		db:      db,
		queries: queries,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.db.Ping(ctx); err != nil {
		JSONError(w, ErrorMsg{Message: "database unavailable"}, http.StatusServiceUnavailable)
		return
	}

	renderJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
