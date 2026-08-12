package users

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserId UserId `json:"user_id"`
	jwt.RegisteredClaims
}

func KeyFunc(*jwt.Token) (any, error) {
	return JWT_SECRET, nil
}

func GetUserId(token *jwt.Token) (UserId, error) {
	return UserId{}, nil
}
