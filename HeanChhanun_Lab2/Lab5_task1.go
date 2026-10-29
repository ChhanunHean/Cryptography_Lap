package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Convert string to binary representation
func toBinary(text string) string {
	var binaryParts []string
	for _, b := range []byte(text) {
		binaryParts = append(binaryParts, fmt.Sprintf("%08b", b))
	}
	return strings.Join(binaryParts, " ")
}

// Convert string to Hexadecimal representation
func toHex(text string) string {
	return hex.EncodeToString([]byte(text))
}

// Convert string to Base64 representation
func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a string: ")
	if scanner.Scan() {
		input := scanner.Text()

		fmt.Println("\n--- Data Representations ---")
		fmt.Printf("Original Text : %s\n", input)
		fmt.Printf("Binary        : %s\n", toBinary(input))
		fmt.Printf("Hexadecimal   : %s\n", toHex(input))
		fmt.Printf("Base64        : %s\n", toBase64(input))
	}
}
