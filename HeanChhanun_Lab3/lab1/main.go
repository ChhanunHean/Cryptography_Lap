package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"hashprogram/lab1/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlistPath := "../nord_vpn.txt"

	// Fallback in case you run from the root folder instead of inside lab1
	if _, err := os.Stat(wordlistPath); os.IsNotExist(err) {
		wordlistPath = "nord_vpn.txt"
	}

	// Open the wordlist file
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Printf("Error opening wordlist file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("======== Password Cracking Program ========")
	fmt.Printf("Target MD5 Hash: %s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	found := false
	attempts := 0

	for scanner.Scan() {
		attempts++
		password := strings.TrimRight(scanner.Text(), "\r\n")

		// Hash the password with our crack package
		currentHash := crack.HashMD5(password)

		// Print verbose output for each attempt
		fmt.Printf("[%d] Testing: %s => Hash: %s\n", attempts, password, currentHash)

		// Check if it matches the target hash
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