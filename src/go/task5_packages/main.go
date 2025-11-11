// File: src/go/task5_packages/main.go
package main
import (
    "fmt"
    "modern-languages-go-basics-Kritika/src/go/task5_packages/greetings"
)


func main() {
    fmt.Println("=== Demonstrating Functions and Packages in Go ===")

    // Call a function from another package
    greetings.GreetUser("Kritika")

    // Call a function that returns a value
    sum := greetings.AddNumbers(10, 20)
    fmt.Printf("The sum of 10 and 20 is: %d\n", sum)

    // Local function example
    result := multiply(5, 6)
    fmt.Printf("Multiplication Result: %d\n", result)
}

// multiply is a local function in main package
func multiply(x int, y int) int {
    return x * y
}
