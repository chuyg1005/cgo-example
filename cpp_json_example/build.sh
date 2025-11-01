#!/bin/bash

# Build script for the JSON processor library

# Create build directory if it doesn't exist
mkdir -p build
cd build

# Configure with CMake
# Note: You need to have vcpkg installed and integrated with CMake
# If vcpkg is installed in ~/vcpkg, use:
# cmake .. -DCMAKE_TOOLCHAIN_FILE=~/vcpkg/scripts/buildsystems/vcpkg.cmake
cmake .. -DCMAKE_TOOLCHAIN_FILE=/path/to/vcpkg/scripts/buildsystems/vcpkg.cmake

# Build the library
cmake --build .

# Copy the library to the project root for easy access
cp libjson_processor.so ../../ || cp libjson_processor.dylib ../../