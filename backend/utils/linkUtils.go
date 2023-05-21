package utils

import (
	"math/rand"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateShortCode(length int) string {

	var shortCode string
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		shortCode += string(alphabet[randomIndex])
	}

	return shortCode
}
