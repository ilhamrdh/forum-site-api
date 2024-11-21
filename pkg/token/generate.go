package token

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/rs/zerolog/log"
)

func GenerateRefreshToken() string {
	b := make([]byte, 18)
	_, err := rand.Read(b)
	if err != nil {
		log.Print("Failed to generate random bytes: ", err)
		return ""
	}
	token := hex.EncodeToString(b)
	log.Print("Generated Refresh Token: ", token)
	return token
}
