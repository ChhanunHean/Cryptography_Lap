package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// HashSHA512 hashes a plain string using SHA-512 and returns the hex string
func HashSHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

// CheckSHA512 returns true if the text's SHA-512 matches the target hash
func CheckSHA512(password string, targetHash string) bool {
	return HashSHA512(password) == targetHash
}