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

	unsortedSlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		unsortedSlice[i] = rand.Intn(max-min) + min
	}

	fmt.Print("\n")

	// (i) Unsorted
	fmt.Print("Unsorted Slice: ", unsortedSlice, "\n")

	// (ii) Sorted in ascending order
	ascendingSlice := make([]int, 10)
	copy(ascendingSlice, unsortedSlice)
	sort.Ints(ascendingSlice)
	fmt.Print("Ascending sequential order: ", ascendingSlice, "\n")

	// (iii) Sorted in descending order
	descendingSlice := make([]int, 10)
	copy(descendingSlice, unsortedSlice)
	sort.Slice(descendingSlice, func(i, j int) bool {
		return descendingSlice[i] >= descendingSlice[j]
	})
	fmt.Print("Descending sequential order: ", descendingSlice, "\n")

	// (iv) Count of odd and even numbers from odd

	fmt.Print("\n")
}
