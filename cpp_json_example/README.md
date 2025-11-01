# C++ JSON Processor with CMake and vcpkg

This example demonstrates using CMake for building a C++ library with vcpkg for package management. The library uses the nlohmann/json third-party library to process JSON data.

## Directory Structure

```
cpp_json_example/
├── include/
│   └── json_processor.h     # Header file with C interface
├── src/
│   └── json_processor.cpp   # Implementation using nlohmann/json
├── CMakeLists.txt           # CMake build configuration
├── vcpkg.json               # vcpkg manifest file
├── build.sh                 # Build script
└── build/                   # Build directory (created during build)
```

## Prerequisites

1. CMake 3.15 or later
2. vcpkg package manager
3. Go 1.16 or later

## Building the C++ Library

1. Make sure vcpkg is installed and integrated with your system
2. Run the build script:
   ```bash
   cd cpp_json_example
   ./build.sh
   ```

## Using the Library with Go

After building the library, you can run the Go example:
```bash
go run cpp_json_example_go.go
```

## Features Demonstrated

- Using CMake with vcpkg for C++ project management
- Integrating third-party libraries (nlohmann/json) via vcpkg
- Creating a C-compatible interface for C++ code
- Calling C++ library functions from Go using CGO
- Proper memory management between C++ and Go
- Error handling in both C++ and Go