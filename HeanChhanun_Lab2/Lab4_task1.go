package main

import (
	"errors"
	"fmt"
)

// Add 
func add(a, b float64) float64 {
	return a + b
}

// Subtract
func sub(a, b float64) float64 {
	return a - b
}

// Multiply
func mul(a, b float64) float64 {
	return a * b
}

// Divide
func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Modulus
func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a % b, nil
}

func main() {
	for {
		fmt.Println("===== Mini Calculator =====")
		fmt.Println("  1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		// Exit option
		if choice == 6 {
			fmt.Println("Goodbye!")
			break
		}

		// Validate choice before requesting numbers
		if choice < 1 || choice > 6 {
			fmt.Println("Invalid choice! Please choose an option between 1 and 6.\n")
			continue
		}

		// Handle Modulo separately as it works on integers
		if choice == 5 {
			var a, b int
			fmt.Print("Enter a: ")
			fmt.Scan(&a)
			fmt.Print("Enter b: ")
			fmt.Scan(&b)

			res, err := mod(a, b)
			if err != nil {
				fmt.Printf("Error: %s\n\n", err)
			} else {
				fmt.Printf("Result: %d %% %d = %d\n\n", a, b, res)
			}
			continue
		}

		// General arithmetic (+, -, *, /) with float support
		var a, b float64
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Printf("Result: %v + %v = %v\n\n", a, b, add(a, b))
		case 2:
			fmt.Printf("Result: %v - %v = %v\n\n", a, b, sub(a, b))
		case 3:
			fmt.Printf("Result: %v * %v = %v\n\n", a, b, mul(a, b))
		case 4:
			res, err := div(a, b)
			if err != nil {
				fmt.Printf("Error: %s\n\n", err)
			} else {
				fmt.Printf("Result: %v / %v = %v\n\n", a, b, res)
			}
		}
	}
}
