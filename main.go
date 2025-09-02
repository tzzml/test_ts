package main

import "fmt"

// sub returns the difference between a and b
func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a sample Go program.")
	
	// Example usage of the sub function
	result := sub(10, 3)
	fmt.Printf("10 - 3 = %d\n", result)
	
	result2 := sub(5, 8)
	fmt.Printf("5 - 8 = %d\n", result2)
}
