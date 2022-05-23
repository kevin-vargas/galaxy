package ctx

import (
	"context"
	"net/http"
)

const correlationIDHeader = "X-Trace-ID"

const (
	contextKeyTraceID   = contextKey("trace_id")
	contextSessionToken = contextKey("session_token")
)

type contextKey string

func (c contextKey) String() string {
	return "request_" + string(c)
}

func GetTraceID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKeyTraceID).(string)
	return id, ok
}

func GetSessionToken(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextSessionToken).(string)
	return id, ok
}

type Option func(context.Context) context.Context

func WithSession(sessionToken string) Option {
	return func(ctx context.Context) context.Context {
		if sessionToken != "" {
			ctx = context.WithValue(ctx, contextSessionToken, sessionToken)
		}
		return ctx
	}
}

func RequestContextBuild(ctx context.Context, r *http.Request, options ...Option) context.Context {
	id := r.Header.Get(correlationIDHeader)
	if id != "" {
		ctx = context.WithValue(ctx, contextKeyTraceID, id)
	}

	for _, option := range options {
		ctx = option(ctx)
	}

	return ctx
}
