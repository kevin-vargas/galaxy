package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/kevin-vargas/galaxy/internal/usecase"

	"net/http"
)

func NewRouter(r *chi.Mux, t usecase.Triangulation, s usecase.Session, m ...func(http.Handler) http.Handler) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	// Routers
	r.Mount("/v1", newRoutes(t, s, m...))
}
