package middlewares

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func extractUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		return uuid.Nil, errors.New("missing token")
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return uuid.Nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(string)
		if !ok {
			return uuid.Nil, errors.New("invalid user ID in token claims")
		}

		userIDUUID, err := uuid.Parse(userID)
		if err != nil {
			return uuid.Nil, err
		}
		return userIDUUID, nil
	}

	return uuid.Nil, errors.New("invalid token claims")
}
