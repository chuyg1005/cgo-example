package main

/*
#cgo LDFLAGS: -L./cpp -lcalculator
#include "./cpp/calculator.h"
*/
import "C"
import "fmt"

func main() {
	// Test all calculator functions
	fmt.Println("Testing C++ Calculator Library")
	
	// Test addition
	result := C.add(5, 3)
	fmt.Printf("5 + 3 = %d (expected: 8)\n", int(result))
	
	// Test subtraction
	result = C.subtract(10, 4)
	fmt.Printf("10 - 4 = %d (expected: 6)\n", int(result))
	
	// Test multiplication
	result = C.multiply(6, 7)
	fmt.Printf("6 * 7 = %d (expected: 42)\n", int(result))
	
	// Test division
	divResult := C.divide(15.0, 3.0)
	fmt.Printf("15.0 / 3.0 = %.2f (expected: 5.00)\n", float64(divResult))
	
	// Test factorial
	result = C.factorial(5)
	fmt.Printf("5! = %d (expected: 120)\n", int(result))
	
	// Test error case
	result = C.factorial(-1)
	fmt.Printf("-1! = %d (expected: -1 for error)\n", int(result))
}