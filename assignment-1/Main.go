package main

import (
	"fmt"
)

func main() {
	assignmentOne()
}

func assignmentOne() {
	var (
		stringOne   string = "This "
		stringTwo   string = "is "
		stringThree string = "a "
		stringFour  string = "concatenated "
		stringFive  string = "string!"
	)

	concatenatedString := stringOne + stringTwo + stringThree + stringFour + stringFive

	fmt.Println("Assignment One:", concatenatedString)
}
