package main

import (
    "fmt"
)

// jiankiewn - Go Implementation
// A simple Hello World program

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to jiankiewn")
    
    // Simple calculation
    numbers := []int{1, 2, 3, 4, 5}
    total := 0
    for _, num := range numbers {
        total += num
    }
    fmt.Printf("Sum of numbers: %d\n", total)
}
