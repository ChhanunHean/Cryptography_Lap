package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// HashMD5 takes a plain string, computes its MD5 hash, and returns the hex string
func HashMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// CheckMD5 returns true if the text's MD5 matches the target hash
func CheckMD5(password string, targetHash string) bool {
	return HashMD5(password) == targetHash
}