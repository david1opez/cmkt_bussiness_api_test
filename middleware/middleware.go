package middleware

import (
	"net/http"
	"os"
)

func Authorize(auth string) int {
	key := os.Getenv("MAGIC_WORD")

	if auth != key {
		return http.StatusUnauthorized
	} else {
		return http.StatusAccepted
	}
}