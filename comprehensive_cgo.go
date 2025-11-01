package main

/*
#include <stdio.h>
#include <stdlib.h>

// A simple C function that adds two integers
int add(int a, int b) {
    return a + b;
}

// A C function that multiplies two integers
int multiply(int a, int b) {
    return a * b;
}

// A C function that prints a message
void print_message(char* msg) {
    printf("C says: %s\n", msg);
}

// A C function that returns a string
char* get_c_string() {
    char* str = malloc(50);
    sprintf(str, "Hello from C function!");
    return str;
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println("=== CGO Examples ===")

    // Example 1: Call the C add function
    a := 10
    b := 20
    result := C.add(C.int(a), C.int(b))
    fmt.Printf("Go calling C add: %d + %d = %d\n", a, b, int(result))

    // Example 2: Call the C multiply function
    x := 5
    y := 6
    product := C.multiply(C.int(x), C.int(y))
    fmt.Printf("Go calling C multiply: %d * %d = %d\n", x, y, int(product))

    // Example 3: Call the C print function
    msg := C.CString("Hello from Go!")
    C.print_message(msg)
    C.free(unsafe.Pointer(msg))

    // Example 4: Get string from C function
    cStr := C.get_c_string()
    defer C.free(unsafe.Pointer(cStr))
    goStr := C.GoString(cStr)
    fmt.Printf("String from C: %s\n", goStr)

    // Example 5: Working with C arrays
    // Create a C array
    cArray := C.malloc(C.size_t(5) * C.sizeof_int)
    defer C.free(cArray)
    
    // Convert to Go slice for easier manipulation
    intArray := (*[1000]C.int)(cArray)[:5:5]
    for i := 0; i < 5; i++ {
        intArray[i] = C.int(i * i) // Fill with squares
    }
    
    fmt.Print("C array values: ")
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", int(intArray[i]))
    }
    fmt.Println()
}