package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

// xorEncrypt
func xorEncrypt(text string, key byte) string {
	bytes := []byte(text)
	result := make([]byte, len(bytes))

	for i := 0; i < len(bytes); i++ {
		result[i] = bytes[i] ^ key
	}

	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read plaintext
	fmt.Print("Enter plaintext: ")
	if !scanner.Scan() {
		return
	}
	plaintext := scanner.Text()

	// Read a single character key
	fmt.Print("Enter single character key: ")
	if !scanner.Scan() {
		return
	}
	keyInput := scanner.Text()

	if len(keyInput) == 0 {
		fmt.Println("Error: Key cannot be empty.")
		return
	}
	key := keyInput[0]

	// 1. Encrypt
	ciphertext := xorEncrypt(plaintext, key)

	// 2. Decrypt
	decrypted := xorEncrypt(ciphertext, key)

	// Display results
	fmt.Println("\n--- XOR Cipher Results ---")
	fmt.Printf("Original Plaintext : %s\n", plaintext)
	fmt.Printf("Key                : '%c' (byte: %d)\n", key, key)
	// Encrypt and Decrypt
	fmt.Printf("Ciphertext (Hex)   : %s\n", hex.EncodeToString([]byte(ciphertext)))
	fmt.Printf("Decrypted Plaintext: %s\n", decrypted)
}
