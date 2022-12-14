package configs

import "os"

func GetJWTSecretKey() string {
	return os.Getenv("JWTSecretKey")
}
