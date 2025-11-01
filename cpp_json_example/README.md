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

## Installing vcpkg

If you don't have vcpkg installed:

```bash
git clone https://github.com/Microsoft/vcpkg.git
cd vcpkg
./bootstrap-vcpkg.sh
cd ..
```

Set the VCPKG_ROOT environment variable:
```bash
export VCPKG_ROOT=/path/to/vcpkg
```

## Building the C++ Library

1. Make sure vcpkg is installed and the VCPKG_ROOT environment variable is set
2. Build the library:
   ```bash
   cd cpp_json_example
   mkdir -p build && cd build
   cmake .. -DCMAKE_TOOLCHAIN_FILE=$VCPKG_ROOT/scripts/buildsystems/vcpkg.cmake
   cmake --build .
   cd ..
   ```
3. Copy the built library to the project root:
   ```bash
   cp build/libjson_processor.* ..
   cd ..
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

## Why This Approach?

1. **vcpkg for Dependency Management**: We use vcpkg to manage the nlohmann/json dependency because:
   - It provides a consistent way to manage C++ dependencies across platforms
   - It handles the complexities of building and linking third-party libraries
   - It integrates well with CMake build system

2. **CMake Build System**: We use CMake because:
   - It's a cross-platform build system that works well with vcpkg
   - It provides good support for C++ projects with external dependencies
   - It generates build files for various compilers and IDEs

3. **Library Placement**: The built library is copied to the project root to:
   - Make it easily discoverable by the Go linker
   - Simplify the linking process in the CGO directives
   - Avoid complex path specifications in the build process