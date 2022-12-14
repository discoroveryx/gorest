package configs

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}
