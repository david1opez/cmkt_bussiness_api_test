package middleware

import (
	"net/http"
	"os"
	"crypto/sha256"
)

func Authorize(auth string) int {
	authE := sha256.New()
	authE.Write([]byte(auth))

	key := os.Getenv("MAGIC_WORD")
	keyE := sha256.New()
	keyE.Write([]byte(key))

	if authE != keyE {
		return http.StatusUnauthorized
	} else {
		return http.StatusAccepted
	}
}