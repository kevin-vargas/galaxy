package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kevin-vargas/galaxy/infra/ctx"
)

func WithCustomContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sessionID string
		c, err := r.Cookie("session_token")

		if err != nil {
			if err == http.ErrNoCookie {
				sessionID = uuid.NewString()
				cookie := &http.Cookie{
					Name:  "session_token",
					Value: sessionID,
				}
				http.SetCookie(w, cookie)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			sessionID = c.Value
		}

		ctxRequest := ctx.RequestContextBuild(r.Context(), r, ctx.WithSession(sessionID))

		r = r.Clone(ctxRequest)
		next.ServeHTTP(w, r)
	})
}
