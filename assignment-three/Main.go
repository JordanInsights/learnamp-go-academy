package main

import (
	"fmt"
)

func main() {
	var i int

	fmt.Print("Enter a number: ")
	digit, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Print(err)
		return
	}

	betweenOneAndTen := digit < 10 && digit > 1

	if betweenOneAndTen {
		fmt.Printf("The entered number: %d is between one and ten\n", digit)
	} else {
		fmt.Printf("The entered number: %d is not between one and ten\n", digit)
	}

}
