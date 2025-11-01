package main

import (
  "fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
  //TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
  // to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
  s := "gopher"
  fmt.Printf("Hello and welcome, %s!\n", s)

  fmt.Println("This project contains CGO examples.")
  fmt.Println("Check out cgo_example.go and comprehensive_cgo.go for basic CGO usage.")
  fmt.Println("For C++ library integration, check out cpp_library_example.go.")
  fmt.Println("For a more complex C++ example with CMake and vcpkg, check out cpp_json_example_go.go.")
  fmt.Println("Run 'go run cgo_example.go' or 'go run comprehensive_cgo.go' for C examples.")
  fmt.Println("Run 'cd cpp && make && go run cpp_library_example.go' for the C++ library example.")
  fmt.Println("See cpp_json_example/README.md for instructions on the CMake/vcpkg example.")
  
  for i := 1; i <= 5; i++ {
	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	fmt.Println("i =", 100/i)
  }
}