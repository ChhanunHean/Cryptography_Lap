package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hashprogram/lab2/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlistPath := "../nord_vpn.txt"

	// Check if wordlist is in current folder or parent folder
	if _, err := os.Stat(wordlistPath); os.IsNotExist(err) {
		wordlistPath = "nord_vpn.txt"
	}

	// 1. Open the wordlist file
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Printf("Error opening wordlist file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("======== SHA-1 Password Cracking Program ========")
	fmt.Printf("Target SHA-1 Hash: %s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	found := false
	attempts := 0

	// 2. Iterate line by line using a for loop
	for scanner.Scan() {
		attempts++
		password := strings.TrimRight(scanner.Text(), "\r\n")

		// 3. Compute SHA-1 hash for this password
		currentHash := crack.HashSHA1(password)

		// 4. Print verbose output
		fmt.Printf("[%d] Testing: %s => Hash: %s\n", attempts, password, currentHash)

		// 5. Compare with target hash
		if currentHash == targetHash {
			fmt.Println("\n===========================================")
			fmt.Printf(" PASSWORD CRACKED SUCCESSFULLY!\n")
			fmt.Printf(" Password: %s\n", password)
			fmt.Printf(" Total Attempts: %d\n", attempts)
			fmt.Println("===========================================")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("\n[-] Password not found in wordlist.")
	}
}