package main

/*
#cgo LDFLAGS: -L./cpp -lcalculator
#include "./cpp/calculator.h"
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println("=== CGO with C++ Library Example ===")

	// Test integer operations
	a := 15
	b := 7

	// Addition
	result := C.add(C.int(a), C.int(b))
	fmt.Printf("C++ library add: %d + %d = %d\n", a, b, int(result))

	// Subtraction
	result = C.subtract(C.int(a), C.int(b))
	fmt.Printf("C++ library subtract: %d - %d = %d\n", a, b, int(result))

	// Multiplication
	result = C.multiply(C.int(a), C.int(b))
	fmt.Printf("C++ library multiply: %d * %d = %d\n", a, b, int(result))

	// Division
	x := 20.0
	y := 3.0
	divResult := C.divide(C.double(x), C.double(y))
	fmt.Printf("C++ library divide: %.2f / %.2f = %.2f\n", x, y, float64(divResult))

	// Factorial
	n := 5
	factResult := C.factorial(C.int(n))
	fmt.Printf("C++ library factorial: %d! = %d\n", n, int(factResult))

	// Test error case for factorial
	n = -3
	factResult = C.factorial(C.int(n))
	if int(factResult) == -1 {
		fmt.Printf("C++ library factorial: %d! = Error (negative number)\n", n)
	} else {
		fmt.Printf("C++ library factorial: %d! = %d\n", n, int(factResult))
	}
}