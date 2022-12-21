package handlers

import (
	"math/rand"
	"strconv"
)

func GenerateVerificationCodeHandler() string {
	newCode := rand.Intn((9999 - 1000) + 1000)

	return strconv.Itoa(newCode)
}
