package handlers

import (
	"fmt"

	"app/auth/configs"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTByUserHandler(userId uint) (string, error) {
	// claims := configs.JWTClaims{
	// 	userId,
	// 	jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 		Issuer:    "test",
	// 	},
	// }

	JWTSecretKey := configs.GetJWTSecretKey()
	fmt.Println("JWTSecretKey", JWTSecretKey)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["UserId"] = userId
	// claims["userName"] = user.Name

	tokenString, err := token.SignedString([]byte(JWTSecretKey))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
