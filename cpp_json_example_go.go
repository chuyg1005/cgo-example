package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L. -ljson_processor -Wl,-rpath,.
#include "cpp_json_example/include/json_processor.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== CGO with C++ Library using nlohmann/json ===")

	// Test JSON processing
	jsonInput := `{
  "name": "Alice",
  "age": 30,
  "city": "New York",
  "hobbies": ["reading", "swimming", "coding"]
}`

	// Convert Go string to C string
	cJsonInput := C.CString(jsonInput)
	defer C.free(unsafe.Pointer(cJsonInput))

	// Call the C++ function
	cResult := C.process_json(cJsonInput)
	if cResult != nil {
		defer C.free_json_result(cResult)

		// Convert C string back to Go string
		goResult := C.GoString(cResult)
		fmt.Println("Processed JSON:")
		fmt.Println(goResult)
	} else {
		fmt.Println("Error: process_json returned NULL")
	}

	// Test with invalid JSON
	invalidJson := `{
  "name": "Bob",
  "age": 25,
  "city": "San Francisco",
  "invalid": json  // Invalid JSON
}`

	cInvalidJson := C.CString(invalidJson)
	defer C.free(unsafe.Pointer(cInvalidJson))

	cErrorResult := C.process_json(cInvalidJson)
	if cErrorResult != nil {
		defer C.free_json_result(cErrorResult)

		goErrorResult := C.GoString(cErrorResult)
		fmt.Println("\nError handling:")
		fmt.Println(goErrorResult)
	}
}
