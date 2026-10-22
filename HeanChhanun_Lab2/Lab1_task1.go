package main // for excute program
import "fmt" // library stand for input and output

func main() {
	var a, b int
	
	fmt.Print("Enter the first integer: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second integer: ")
	fmt.Scan(&b)

	fmt.Printf("\nInitial Values : a = %d, b = %d\n\n", a, b)

	num := a
	fmt.Printf("num := a     -> num = %d\n", num)

	// add
	num += b
	fmt.Printf("num += b     -> num = %d\n", num)

	// subtract
	num -= b 
	fmt.Printf("num -= b     -> num = %d\n", num)

	// multiply
	num *= b 
	fmt.Printf("num *= b     -> num = %d\n", num)

	//divice
	if b != 0 {
		num /= b
		fmt.Printf("num /= b     -> num %d\n", num)
	}else {
		fmt.Println("res /=     -> Can't divide by zero")
	}

	// modulus
	if b != 0 {
		num %= b
		fmt.Printf("num %%= b     -> num = %d\n", num)
	}else{
		fmt.Println("res /= b    -> Can't module by zero")
	}
}