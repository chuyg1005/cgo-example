package main

/*
#include <stdio.h>
#include <stdlib.h>

// A simple C function that adds two integers
int add(int a, int b) {
    return a + b;
}

// A C function that prints a message
void print_message(char* msg) {
    printf("C says: %s\n", msg);
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    // Call the C add function
    a := 10
    b := 20
    result := C.add(C.int(a), C.int(b))
    fmt.Printf("Go calling C add: %d + %d = %d\n", a, b, int(result))

    // Call the C print function
    msg := C.CString("Hello from Go!")
    C.print_message(msg)
    C.free(unsafe.Pointer(msg))
}