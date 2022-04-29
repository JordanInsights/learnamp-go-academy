package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Provide seed
	rand.Seed(time.Now().Unix())

	min := 1
	max := 11

	var collection [10]int

	i := 0
	for i < 10 {
		collection[i] = rand.Intn(max-min) + min
		i += 1
	}

	fmt.Print("\n")

	fmt.Print("Unsorted Array: ", collection, "\n")

	sort.Ints(collection[:])
	fmt.Print("Ascending sequential order: ", collection, "\n")

	sort.Slice(collection[:], func(i, j int) bool {
		return collection[i] >= collection[j]
	})
	fmt.Print("Descending sequential order: ", collection, "\n")

	fmt.Print("\n")
}
