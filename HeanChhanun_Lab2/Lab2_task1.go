package main

import (
	"fmt"
)

func main() {
	var a, b int

	// user input
	fmt.Print("Enter first integer: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Print("Enter second integer: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println("\n--- Evaluating Conditions ---")

	// Both positive using &&
	bothPositive := (a > 0) && (b > 0)
	fmt.Printf("Both are positive (&&): %t\n", bothPositive)

	// One greater than the other using ||
	oneGreaterThanOther := (a > b) || (b > a)
	fmt.Printf("One is greater than the other (||): %t\n", oneGreaterThanOther)

	// Not equal using !
	notEqual := !(a == b)
	fmt.Printf("Not equal (!): %t\n", notEqual)
}
