package main
import "fmt"

func main() {
    // if-else
    age := 21
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }

    // switch
    grade := "A"
    switch grade {
    case "A":
        fmt.Println("Excellent!")
    case "B":
        fmt.Println("Good job!")
    default:
        fmt.Println("Keep improving!")
    }

    // for loop
    fmt.Println("Numbers from 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // range loop
    nums := []int{10, 20, 30, 40}
    for index, value := range nums {
        fmt.Printf("Index %d: %d\n", index, value)
    }
}
