package main
import "fmt"

// Bitwise AND
func myAND(a, b int) {
	fmt.Printf("Bitwise AND (%d & %d): %d (Binary: %08b)\n", a, b, a&b, a&b)
}
// Bitwise OR
func myOR(a, b int) {
	fmt.Printf("Bitwise OR  (%d | %d): %d (Binary: %08b)\n", a, b, a|b, a|b)
}
// Bitwise XOR
func myXor(a, b int) {
	fmt.Printf("Bitwise XOR (%d ^ %d): %d (Binary: %08b)\n", a, b, a^b, a^b)
}
// Bitwise NOT (unary ^ in Go)
func myNOT(a, b int) {
	fmt.Printf("Bitwise NOT (^%d): %d\n", a, ^a)
	fmt.Printf("Bitwise NOT (^%d): %d\n", b, ^b)
}
// Shift operations (using b as shift amount)
func myShifts(a int, shift uint) {
	fmt.Printf("Left Shift  (%d << %d): %d (Binary: %08b)\n", a, shift, a<<shift, a<<shift)
	fmt.Printf("Right Shift (%d >> %d): %d (Binary: %08b)\n", a, shift, a>>shift, a>>shift)
}
// Demonstration of assignment operators
func demonstrateAssignments(initial int) {
	fmt.Println("\n--- Assignment Operators Demo ---")
	val := initial
	fmt.Printf("Initial value: %d\n", val)

	val += 5 // val = val + 5
	fmt.Printf("After val += 5  : %d\n", val)

	val -= 3 // val = val - 3
	fmt.Printf("After val -= 3  : %d\n", val)

	val *= 2 // val = val * 2
	fmt.Printf("After val *= 2  : %d\n", val)

	val &= 7 // Bitwise AND assignment: val = val & 7
	fmt.Printf("After val &= 7  : %d\n", val)

	val |= 2 // Bitwise OR assignment: val = val | 2
	fmt.Printf("After val |= 2  : %d\n", val)

	val ^= 1 // Bitwise XOR assignment: val = val ^ 1
	fmt.Printf("After val ^= 1  : %d\n", val)

	val <<= 2 // Left shift assignment: val = val << 2
	fmt.Printf("After val <<= 2 : %d\n", val)

	val >>= 1 // Right shift assignment: val = val >> 1
	fmt.Printf("After val >>= 1 : %d\n", val)
}

func main() {
	a := 12 // 00001100 in binary
	b := 5  // 00000101 in binary
	fmt.Printf("Values: a = %d (%08b), b = %d (%08b)\n\n", a, a, b, b)
	// Functions
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	// Shift operations
	fmt.Println()
	myShifts(a, 2)
	// demo
	demonstrateAssignments(10)
}
