package main

/*
// Note: These directives would be used when the library is built
//#cgo LDFLAGS: -L./cpp_json_example -ljson_processor
//#include "./cpp_json_example/include/json_processor.h"
//#include <stdlib.h>
*/
import "C"
import "fmt"

// This is a placeholder test that shows how the code would be used
// In a real environment with vcpkg and CMake properly set up, this would work
func main() {
	fmt.Println("=== CGO with C++ Library using nlohmann/json ===")
	fmt.Println("This example requires vcpkg and CMake to build the library.")
	fmt.Println("See cpp_json_example/README.md for build instructions.")
	
	// In a real implementation, we would do:
	/*
	jsonInput := `{"name": "Alice", "age": 30}`
	cJsonInput := C.CString(jsonInput)
	defer C.free(unsafe.Pointer(cJsonInput))
	
	cResult := C.process_json(cJsonInput)
	if cResult != nil {
		defer C.free_json_result(cResult)
		goResult := C.GoString(cResult)
		fmt.Println(goResult)
	}
	*/
	
	fmt.Println("\nExample JSON input:")
	fmt.Println(`{"name": "Alice", "age": 30, "city": "New York", "hobbies": ["reading", "swimming", "coding"]}`)
	
	fmt.Println("\nExpected output after processing:")
	fmt.Println(`{
  "age": 30,
  "city": "New York",
  "greeting": "Hello, Alice!",
  "hobbies": [
    "reading",
    "swimming",
    "coding"
  ],
  "name": "Alice",
  "processed": true,
  "processed_at": 1234567890
}`)
}