package middleware

import (
	"context"
	"math/rand"
	"net/http"
)

type requestIDMiddleware struct {
	next http.Handler
}

type requestIDContextKey string

const (
	requestIDKey = requestIDContextKey("requestID")
)

func NewRequestID(next http.Handler) http.Handler {
	return &requestIDMiddleware{
		next: next,
	}
}

func (mw *requestIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := rand.Int()

	ctx := r.Context()
	ctx = context.WithValue(ctx, requestIDKey, id)
	r = r.WithContext(ctx)
	mw.next.ServeHTTP(w, r)
}

func RequestIDFromContext(ctx context.Context) int {
	return ctx.Value(requestIDKey).(int)
}
