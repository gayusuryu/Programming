// switch case example

package main

import (
	"fmt"
)

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("many")
	}

	// switch without an expression
	switch {
	case x > 100:
		fmt.Println("x is very Big")
	case x > 10:
		fmt.Println("x is Big")
	default:
		fmt.Println("x is Small")

	}
}
