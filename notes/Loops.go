package main

import (
	"fmt"
)

// for is Go's ONLY LOOPING CONSTRUCT
// There are a couple of basic types

func Loops() {
	// The most basic types of loop, with a single condition
	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/conditions/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition; will loops until it hits a break
	// or a return from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// continue can also be used to initiate the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
