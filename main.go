package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a sample Go program.")
	
	// 示例：使用add函数
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
}
