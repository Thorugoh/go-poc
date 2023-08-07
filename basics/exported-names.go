package main

import (
	"fmt"
	"math"
)

// A name is exported if it begins with a capital letter.
// When importing a package we can refer only to its exported names.

func main() {
	fmt.Println(math.Pi)
}
