cmake_minimum_required(VERSION 3.15)
project(JsonProcessor LANGUAGES CXX)

# Set C++ standard
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find vcpkg packages
find_package(nlohmann_json CONFIG REQUIRED)

# Create the library
add_library(json_processor SHARED
    src/json_processor.cpp
)

# Set library properties
set_target_properties(json_processor PROPERTIES
    PUBLIC_HEADER include/json_processor.h
    INSTALL_RPATH "@loader_path"
    BUILD_WITH_INSTALL_RPATH TRUE
)

# Link with nlohmann_json
target_link_libraries(json_processor PRIVATE nlohmann_json::nlohmann_json)

# Include directories
target_include_directories(json_processor PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}/include
)

# Install rules
install(TARGETS json_processor
    LIBRARY DESTINATION lib
    PUBLIC_HEADER DESTINATION include
)