package delivery

import "net/http"

func GeneralRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health,)

	return mux
}

func AppRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", GeneralRouter(handler))

	return mux
}