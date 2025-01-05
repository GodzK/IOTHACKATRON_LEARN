package main

import "fmt"

func main() {
    // Basic string formatting with Sprintf
    name := "Gopher"
    age := 25

    // Create a formatted string without printing
    greeting := fmt.Sprintf("Hello, my name is %s", name)
    
    // Multiple variable formatting
    personInfo := fmt.Sprintf("Name: %s, Age: %d", name, age)
    
    // Formatting with different types
    price := 67.1258
    productInfo := fmt.Sprintf("Product costs $%.2f", price)
    
    // More complex formatting
    status := true
    userStatus := fmt.Sprintf("Is active: %t", status)

    // Print formatted strings
    fmt.Println(greeting)
    fmt.Println(personInfo)
    fmt.Println(productInfo)
    fmt.Println(userStatus)
}