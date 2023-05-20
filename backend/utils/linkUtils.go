package utils

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateShortLink(length int) string {

	var shortLink string
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		shortLink += string(alphabet[randomIndex])
	}

	return fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), shortLink)
}
