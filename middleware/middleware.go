package middleware

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
)

func Authorize(auth string) int {
	authE := sha256.New()
	authE.Write([]byte(auth))

	key := os.Getenv("MAGIC_WORD")
	keyE := sha256.New()
	keyE.Write([]byte(key))

	fmt.Println(key)
	fmt.Println(keyE.Sum(nil))

	fmt.Println(auth)
	fmt.Println(authE.Sum(nil))

	if authE != keyE {
		return http.StatusUnauthorized
	} else {
		return http.StatusAccepted
	}
}