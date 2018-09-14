package printer

import "fmt"

// SayHi says hello with a hardcoded version that should be changed whenever the version is incremented
func SayHi() {
	fmt.Println("Hi from go-module-test! This is version 2")
}
