package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hashprogram/lab3/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	wordlistPath := "../nord_vpn.txt"

	// Check if wordlist is in current folder 
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

	fmt.Println("======== SHA-512 Password Cracking Program ========")
	fmt.Printf("Target SHA-512 Hash:\n%s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	found := false
	attempts := 0

	// 2. Iterate line by line using a for loop
	for scanner.Scan() {
		attempts++
		password := strings.TrimRight(scanner.Text(), "\r\n")

		// 3. Compute SHA-512 hash for this password
		currentHash := crack.HashSHA512(password)

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