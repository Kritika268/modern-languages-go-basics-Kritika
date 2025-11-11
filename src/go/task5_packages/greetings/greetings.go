// File: src/go/task5_packages/greetings/greetings.go
package greetings

import "fmt"

// GreetUser prints a simple greeting message
func GreetUser(name string) {
    fmt.Printf("Hello, %s! Welcome to Go Packages.\n", name)
}

// AddNumbers returns the sum of two integers
func AddNumbers(a int, b int) int {
    return a + b
}

// sayGoodbye is an unexported function (not accessible outside)
func sayGoodbye(name string) {
    fmt.Printf("Goodbye, %s!\n", name)
}
