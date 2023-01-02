package handlers

import (
	"app/auth/configs"
	"app/auth/exceptions"

	"github.com/golang-jwt/jwt/v4"
)

func ParseJWTHandler(bearerValue string) (*configs.JWTClaims, error) {

	JWTSecretKey := configs.GetJWTSecretKey()

	// token, _ := jwt.ParseWithClaims(bearerValue, &configs.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	hmacSampleSecret := []byte(JWTSecretKey)
	// 	return hmacSampleSecret, nil
	// })

	token, _ := jwt.ParseWithClaims(bearerValue, &configs.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})


	// if err =! nil {
	// 	return nil, exceptions.JWTFailedError
	// }


	if token == nil {
		return nil, exceptions.JWTFailedError
	}


	if token.Valid == false {
		return nil, exceptions.JWTFailedError
	}

	claims, ok := token.Claims.(*configs.JWTClaims)

	if ok {
		return claims, nil
	}

	// {
	// 	return claims, nil
	// }
	// token, _ := jwt.Parse(bearerValue, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// })

	// if token.Valid {
	// 	return token, nil
	// }

	return nil, exceptions.JWTFailedError
}
