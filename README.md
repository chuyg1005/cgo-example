# Learn CGO

This project contains examples of using CGO to call C/C++ code from Go.

## Files

1. `cgo_example.go` - A simple example showing basic CGO usage with C
2. `comprehensive_cgo.go` - A more comprehensive example showing various CGO features with C
3. `cpp/` - Directory containing C++ library files:
   - `calculator.h` and `calculator.cpp` - C++ library implementation
   - `Makefile` - Build script for the C++ library
4. `cpp_library_example.go` - Example showing how to call the C++ library from Go
5. `test_calculator.go` - Test file for the C++ calculator library
6. `cpp_json_example/` - Directory containing a more complex C++ example with CMake and vcpkg:
   - Uses nlohmann/json third-party library
   - Built with CMake and vcpkg package management
   - See `cpp_json_example/README.md` for details
7. `cpp_json_example_go.go` - Go file that calls the C++ JSON processor library

## Running the Examples

To run the simple C example:
```bash
go run cgo_example.go
```

To run the comprehensive C example:
```bash
go run comprehensive_cgo.go
```

To build and run the C++ library example:
```bash
# Build the C++ library
cd cpp && make

# Run the Go program that uses the C++ library
go run cpp_library_example.go
```

## CGO Features Demonstrated

### C Examples:
- Calling C functions from Go
- Passing data between Go and C
- Converting Go strings to C strings
- Getting strings from C functions
- Working with C arrays
- Memory management with C.malloc and C.free

### C++ Library Example:
- Creating a C++ shared library
- Using extern "C" for C-compatible function interfaces
- Linking a C++ library with a Go program using CGO
- Calling C++ library functions from Go

## Requirements

- Go 1.16 or later
- A C/C++ compiler (usually GCC or Clang)
- Make build tool