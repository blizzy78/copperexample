package middleware

import (
	"context"
	"math/rand"
	"net/http"
)

type requestIDContextKey string

const requestIDKey = requestIDContextKey("requestID")

func NewRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// generate some random fake ID
		id := rand.Int()

		// put ID into request's context
		ctx := r.Context()
		ctx = withRequestID(ctx, id)
		r = r.WithContext(ctx)

		// call next handler
		next.ServeHTTP(w, r)
	})
}

func withRequestID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFromContext(ctx context.Context) int {
	return ctx.Value(requestIDKey).(int)
}
