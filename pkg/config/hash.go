package config

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	salt     = "qwerty"
	TokenTTL = 12 * time.Hour
	SignKey  = "98fhjvpdfh4r398hv"
)

func GenerateHash(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
