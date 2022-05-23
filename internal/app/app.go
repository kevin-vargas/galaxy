package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevin-vargas/galaxy/config"
	v1 "github.com/kevin-vargas/galaxy/internal/controller/v1"
	"github.com/kevin-vargas/galaxy/internal/usecase"
	"github.com/kevin-vargas/galaxy/internal/usecase/finder"
	"github.com/kevin-vargas/galaxy/internal/usecase/satellites"
	"github.com/kevin-vargas/galaxy/internal/usecase/session"
)

func Run(cfg *config.Config) {
	r := chi.NewRouter()

	s := satellites.New()
	f := finder.New()
	t := usecase.NewTriangulation(s, f)
	ses := session.New()

	v1.NewRouter(r, t, ses)

	err := http.ListenAndServe(":"+cfg.App.Port, r)

	if err != nil {
		panic(err)
	}
}
