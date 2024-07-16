package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type ContextKey string

type ContextKeys struct {
	UserIDKey     ContextKey
	IsVerifiedKey ContextKey
}

var ContextKeyNames = ContextKeys{
	UserIDKey:     "userID",
	IsVerifiedKey: "isVerified",
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		splitAuthHeader := strings.Split(authHeader, " ")
		if len(splitAuthHeader) != 2 {
			http.Error(w, "Malformed Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := splitAuthHeader[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims[string(ContextKeyNames.UserIDKey)].(string)
			if !ok {
				http.Error(w, "Invalid user ID in token claims", http.StatusUnauthorized)
				return
			}

			isVerified, ok := claims[string(ContextKeyNames.IsVerifiedKey)].(bool)
			if !ok {
				http.Error(w, "Invalid verification status in token claims", http.StatusUnauthorized)
				return
			}

			if !isVerified {
				http.Error(w, "User is not verified", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKeyNames.UserIDKey, userID)
			ctx = context.WithValue(ctx, ContextKeyNames.IsVerifiedKey, isVerified)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

	})
}

func ExtractUserIDFromContext(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	userID, ok := r.Context().Value(ContextKeyNames.UserIDKey).(string)
	if !ok {
		return uuid.Nil, errors.New("failed to extract user id")
	}
	// Convert the user ID to uuid.UUID
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, err
	}

	return userIDUUID, nil
}
