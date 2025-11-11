package main
import "fmt"

func main() {
    // Array
    var arr = [3]string{"Go", "Java", "Kotlin"}
    fmt.Println("Array:", arr)

    // Slice
    slice := []int{1, 2, 3}
    slice = append(slice, 4, 5)
    fmt.Println("Slice after append:", slice)

    // Copy
    copySlice := make([]int, len(slice))
    copy(copySlice, slice)
    fmt.Println("Copied Slice:", copySlice)

    // Map
    student := map[string]int{"Kritika": 21, "Riya": 20}
    student["Mehul"] = 22
    fmt.Println("Map:", student)
}
