package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Provide seed
	rand.Seed(time.Now().Unix())

	min := 1
	max := 10

	var collection [10]int

	i := 0
	for i < 10 {
		collection[i] = rand.Intn(max-min) + min
		i += 1
	}

	fmt.Print(collection)
}
