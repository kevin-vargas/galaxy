package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevin-vargas/galaxy/infra/ctx"
	"github.com/kevin-vargas/galaxy/infra/middleware"
	"github.com/kevin-vargas/galaxy/internal/entity"
	"github.com/kevin-vargas/galaxy/internal/usecase"
)

type routes struct {
	t usecase.Triangulation
	s usecase.Session
}

func newRoutes(t usecase.Triangulation, s usecase.Session, m ...func(http.Handler) http.Handler) http.Handler {
	router := chi.NewRouter()
	router.Use(m...)
	router.Use(middleware.WithCustomContext)

	r := &routes{t, s}

	router.Post("/topsecret", r.handleTopSecret())
	router.Post("/topscret-split/{satellite_name}", r.handleTopSecretSplit())
	router.Get("/topscret-split", r.handleTopSecretSplitGet())

	return router
}

func (r *routes) getSecret(w http.ResponseWriter, satellites entity.SatelliteElems) {
	res := newResponse(w)
	positions := satellites.Positions()
	points, err := r.t.GetPoints(positions...)
	if err != nil {
		res.withHTTPError(http.StatusConflict, "get positions", IncorrectBody(err.Error()))
		return
	}
	messages := satellites.Messages()
	message, err := r.t.GetMessage(messages...)
	if err != nil {
		res.withHTTPError(http.StatusConflict, "get message", IncorrectBody(err.Error()))
		return
	}

	res.
		withStatusCode(http.StatusOK).
		withJSON(topsecretResponse{
			Positions: points,
			Message:   message,
		})
}

func (r *routes) handleTopSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var satellites entity.Satellites

		if ok := validateRequest(w, req, &satellites); ok {
			r.getSecret(w, satellites.SatelliteElems)
		}
	}
}

func validateRequest(w http.ResponseWriter, r *http.Request, elem any) (ok bool) {
	const BaseMsgError = "On Request"
	res := newResponse(w)

	headerContentTtype := r.Header.Get(headerContentType)
	if headerContentTtype != contentType {
		res.
			withHTTPError(http.StatusUnsupportedMediaType,
				BaseMsgError,
				IncorrectHeader(headerContentType, "Content Type is not "+contentType))
		return
	}

	err := decodeReadCloser(r.Body, elem)
	if err != nil {
		res.withHTTPError(
			http.StatusBadRequest,
			BaseMsgError,
			IncorrectBody(err.Error()),
		)
		return
	}

	return true
}

func (r *routes) handleTopSecretSplit() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		satelliteName := chi.URLParam(req, "satellite_name")
		var elem entity.SatelliteElem

		if ok := validateRequest(w, req, &elem); ok {
			elem.SatelliteName = satelliteName
			if sessionID, ok := ctx.GetSessionToken(req.Context()); ok {
				if satellites, ok := r.s.Get(sessionID); ok {
					satellites = append(satellites, &elem)
					r.s.Set(sessionID, satellites)
				} else {
					r.s.Set(sessionID, []*entity.SatelliteElem{&elem})
				}
				w.WriteHeader(http.StatusCreated)
				return
			}
		}
	}
}

func (r *routes) handleTopSecretSplitGet() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if sessionID, ok := ctx.GetSessionToken(req.Context()); ok {
			if satellites, ok := r.s.Get(sessionID); ok {
				r.getSecret(w, satellites)
				return
			}
		}
	}
}

func decodeReadCloser(rc io.ReadCloser, elem interface{}) error {
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(rc)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(elem)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return errors.New("wrong Type provided for field " + unmarshalErr.Field)
		}
	}
	return err
}

type topsecretResponse struct {
	Positions []*entity.Point `json:"positions"`
	Message   string          `json:"message"`
}
