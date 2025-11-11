package main
import "fmt"

func updateValue(x *int) {
    *x = *x + 10
}

func main() {
    num := 5
    ptr := &num

    fmt.Println("Before update:", num)
    updateValue(ptr)
    fmt.Println("After update:", num)
    fmt.Println("Pointer value:", *ptr)

    // Garbage collection note
    fmt.Println("\nGo automatically performs garbage collection — freeing unused memory.")
}
