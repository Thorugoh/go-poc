package main

import "fmt"

// Go return values may be named. If so, they are treated as variables defined at the top of the function.
// Names are used to document the meaning of return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments returns the named return values.
	// This is a "naked" return
	return
}

func main() {
	x, y := split(17)
	fmt.Println(x, y)
}
