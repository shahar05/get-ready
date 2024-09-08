package middleware

import (
	"context"
	"net/http"
	"phonebook-api/token"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get(authorizationHeaderKey)

			// Check header
			if len(authorizationHeader) == 0 {
				http.Error(w, "Authorization header is not provided", http.StatusUnauthorized)
				return
			}

			// Check 2 fields
			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			// Check bearer
			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				http.Error(w, "Unsupported authorization type", http.StatusUnauthorized)
				return
			}

			// Verify Token
			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// Set payload in context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "payload", payload)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
