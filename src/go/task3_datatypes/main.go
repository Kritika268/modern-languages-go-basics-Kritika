package main
import "fmt"

func main() {
    // Variable declarations
    var name string = "Kritika"
    age := 20 // shorthand declaration
    const college string = "K.R. Mangalam University"

    // Data types
    var isStudent bool = true
    var cgpa float32 = 8.9
    var year int = 2025

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("College:", college)
    fmt.Println("CGPA:", cgpa)
    fmt.Println("Graduation Year:", year)
    fmt.Println("Is Student:", isStudent)

    // Type conversion
    var intToFloat = float64(year)
    fmt.Println("Converted year to float:", intToFloat)
}
