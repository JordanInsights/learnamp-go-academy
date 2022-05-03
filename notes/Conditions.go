package main

import (
	"fmt"
)

func Conditions() {
	// Conditions work basically the same as in JavaScript

	if 7%2 == 0 {
		fmt.Println("Seven is even")
	} else {
		fmt.Println("Seven is odd")
	}

	// Some weirdness though; apparently statements can precede conditionals
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
