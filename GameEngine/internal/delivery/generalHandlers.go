package delivery

import (
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct{
	db *pgxpool.Pool
}

func NewHandler(
	db *pgxpool.Pool,
) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.db.Ping(ctx); err != nil {
		JSONError(w, ErrorMsg {Message: "database unavailable"}, http.StatusServiceUnavailable)
		return
	}

	renderJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}