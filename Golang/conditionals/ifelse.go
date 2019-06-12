// Example of if statement

package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is Big")
	}

	if x > 100 {
		fmt.Println("x is very Big")
	} else {
		fmt.Println("x is not that Big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just Right")
	}

	if x < 20 || x < 30 {
		fmt.Println("x is out of Range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}
