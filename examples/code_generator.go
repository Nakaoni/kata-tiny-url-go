package katatinyurl

import (
	"math/rand/v2"
	"strings"
	// "strings"
)

const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CODE_LENGTH = 6
)

// Generates and returns a random 6 char-length code
func GenerateCode() string {
	var sb strings.Builder
	sb.Grow(CODE_LENGTH)

	for range CODE_LENGTH {
		randomIndex := rand.IntN(len(CHARSET))
		sb.WriteByte(CHARSET[randomIndex])
	}

	return sb.String()
}
