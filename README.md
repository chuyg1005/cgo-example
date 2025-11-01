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

To build and run the C++ JSON processor example with vcpkg:
```bash
# Install vcpkg if not already installed
git clone https://github.com/Microsoft/vcpkg.git
cd vcpkg
./bootstrap-vcpkg.sh
cd ..

# Set VCPKG_ROOT environment variable
export VCPKG_ROOT=/path/to/vcpkg

# Build the C++ JSON processor library with vcpkg
cd cpp_json_example
mkdir -p build && cd build
cmake .. -DCMAKE_TOOLCHAIN_FILE=$VCPKG_ROOT/scripts/buildsystems/vcpkg.cmake
cmake --build .
cd ../..

# Run the Go program that uses the C++ JSON processor library
# No need to copy the library - it's found automatically through CGO directives
go run cpp_json_example_go.go
```

## CGO Features Demonstrated

### C Examples:
- Calling C functions from Go
- Passing data between Go and C
- Converting Go strings to C strings
- Getting strings from C functions
- Working with C arrays
- Memory management with C.malloc and C.free

### C++ Library Examples:
- Creating a C++ shared library
- Using extern "C" for C-compatible function interfaces
- Linking a C++ library with a Go program using CGO
- Calling C++ library functions from Go
- Using CMake with vcpkg for C++ project management
- Integrating third-party libraries (nlohmann/json) via vcpkg
- Proper memory management between C++ and Go
- Error handling in both C++ and Go

## Requirements

- Go 1.16 or later
- A C/C++ compiler (usually GCC or Clang)
- Make build tool
- CMake 3.15 or later (for JSON example)
- vcpkg package manager (for JSON example)

## Why This Approach?

1. **CGO Linking**: The JSON example uses specific CGO directives to link with the C++ library:
   - `#cgo CXXFLAGS: -std=c++11` specifies the C++ standard
   - `#cgo LDFLAGS: -L${SRCDIR}/cpp_json_example/build -ljson_processor -Wl,-rpath,${SRCDIR}/cpp_json_example/build` specifies:
     - `-L${SRCDIR}/cpp_json_example/build` to look for libraries in the build directory
     - `-ljson_processor` to link with the json_processor library
     - `-Wl,-rpath,${SRCDIR}/cpp_json_example/build` to set the runtime library search path
   - Using `${SRCDIR}` ensures the paths are relative to the source file location

2. **vcpkg for Dependency Management**: We use vcpkg to manage the nlohmann/json dependency because:
   - It provides a consistent way to manage C++ dependencies across platforms
   - It handles the complexities of building and linking third-party libraries
   - It integrates well with CMake build system

3. **CMake Build System**: We use CMake because:
   - It's a cross-platform build system that works well with vcpkg
   - It provides good support for C++ projects with external dependencies
   - It generates build files for various compilers and IDEs

4. **No Library Copying Needed**: We've eliminated the need to copy the built library to the project root by:
   - Using `${SRCDIR}` in the CGO directives to specify paths relative to the source file
   - Setting the rpath during linking to ensure the library can be found at runtime
   - Configuring the CMake build to set appropriate rpath values in the library itself