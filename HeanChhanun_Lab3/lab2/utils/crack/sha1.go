package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// HashSHA1 hashes a plain string using SHA-1 and returns the hex string
func HashSHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// CheckSHA1 returns true if the text's SHA-1 matches the target hash
func CheckSHA1(password string, targetHash string) bool {
	return HashSHA1(password) == targetHash
}