package main

import (
	"fmt"
	"strconv"
)

func DeclaringVariables() {
	// Syntax to use when you're declaring a variable
	// that you're not using immediately
	var i int
	i = 1
	fmt.Printf("%v, %T\n", i, i)

	// This is the syntax to use when you need specific
	// control over the variable type, for example if
	// you need a type of float32 which cannot be declared
	// implicitly
	var j float32 = 2
	fmt.Printf("%v, %T\n", j, j)

	// This is the most concise syntax for declaring a
	// a variable with a type
	k := 3.
	fmt.Printf("%v, %T\n", k, k)

	// Declare multiple variables at once
	var (
		l int     = 4
		m float64 = 5
	)
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", m, m)

	// Convert integer to string
	var n int = 6
	var o string
	o = strconv.Itoa(n)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", o, o)
}
