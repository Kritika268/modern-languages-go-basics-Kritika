package main
import "fmt"

// Struct
type Student struct {
    Name  string
    Age   int
    Grade string
}

// Method
func (s Student) displayInfo() {
    fmt.Printf("Name: %s | Age: %d | Grade: %s\n", s.Name, s.Age, s.Grade)
}

func main() {
    s1 := Student{Name: "Kritika", Age: 21, Grade: "A+"}
    s1.displayInfo()
}
