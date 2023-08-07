package main

import "fmt"

// we can omit type of consecutive parameters if they are the same (must type the last one)
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
