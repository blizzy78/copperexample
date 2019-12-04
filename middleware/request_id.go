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

const requestIDKey requestIDContextKey = "requestID"

func NewRequestID(next http.Handler) http.Handler {
	return &requestIDMiddleware{
		next: next,
	}
}

func (mw *requestIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// generate some random fake ID
	id := rand.Int()

	// put ID into request's context
	ctx := withRequestID(r.Context(), id)
	r = r.WithContext(ctx)

	// call next handler
	mw.next.ServeHTTP(w, r)
}

func withRequestID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFromContext(ctx context.Context) int {
	return ctx.Value(requestIDKey).(int)
}
