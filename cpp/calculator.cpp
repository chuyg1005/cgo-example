#include "calculator.h"

// Add two integers
int add(int a, int b) {
    return a + b;
}

// Subtract two integers
int subtract(int a, int b) {
    return a - b;
}

// Multiply two integers
int multiply(int a, int b) {
    return a * b;
}

// Divide two doubles
double divide(double a, double b) {
    if (b == 0) {
        return 0; // Handle division by zero
    }
    return a / b;
}

// Calculate factorial of a number
int factorial(int n) {
    if (n < 0) {
        return -1; // Error case
    }
    if (n == 0 || n == 1) {
        return 1;
    }
    int result = 1;
    for (int i = 2; i <= n; i++) {
        result *= i;
    }
    return result;
}