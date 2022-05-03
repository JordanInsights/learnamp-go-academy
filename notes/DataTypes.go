package main

import (
	"fmt"
)

func DataTypes() {
	// Boolean Data
	// Basic declaration
	var m bool = true
	n := false

	// As a result of a logical test
	o := 1 == 1
	p := 1 == 2

	// When you initialise a variable in Go without giving it a value
	// it has a zero value, in this case the zero value for a boolean is false
	var q bool

	fmt.Println(m, n, o, p, q)

	// Number Data
	r := 42
	fmt.Println(r)

	// Go wont handle implicit type conversion
	// need to convert integers of different types
	// in order to sum them
	var s int = 1
	var t int8 = 1
	fmt.Println(s + int(t))

	// bit operators
	// These look at what bits are set in the numbers

	u := 10 // 1010 in binary
	v := 3  // 0011

	// AND Operator
	fmt.Println(u & v) // if the bit is truthy (1) in both numbers it makes it through and therefore results in:
	// 0010, which equals 2

	// OR Operator
	fmt.Println(u | v) // if the bit is truthy in u OR v it makes it through and therefore results in:
	// 1011, which equals 11

	// EXCLUSIVE OR Operator
	fmt.Println(u ^ v) // checks the bit is exclusive, i.e. different between the two; or for example 1 in u and 0 in v will make it through
	// 1001, which equals 9

	// AND NOT operator
	fmt.Println(u &^ v) // opposite of OR, will make it through if NEITHER has the bit set
	// 0100, which equals 8

	// bit shifting
	x := 8 // 2^3

	fmt.Println(x << 3) // bit shift x left three places
	// s^3 * 2^3 = 2^6 = 64

	fmt.Println(x >> 3) // bit shift x right three places
	// 2^3 / 2^3 = 2^0 = 1

	// Floating Point numbers
	// all of these are valid declariation syntax
	// NO BIT OPERATORS OR BIT SHIFTING
	y := 3.14
	y = 13.7e72

	var z float32 = 2.1e14

	// 2.1e14 = 2.1 * 10^14

	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

	// Complex numbers
	var a complex128 = 1 + 2i
	b := 2 + 5.2i
	c := a + b

	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	var d complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", d, d)

	// Text type
	// Strings - and UTF character
	stringLiteral := "This is a string"
	fmt.Printf("%v, %T\n", stringLiteral, stringLiteral)
	fmt.Printf("%v, %T\n", stringLiteral[2], stringLiteral[2]) // strings are aliases for bytes
	fmt.Printf("%v, %T\n", string(stringLiteral[2]), string(stringLiteral[2]))

	// Rune type
	// Rune represents a UTF32 character
	// UTF32 is a bit weird, a character in UTF32 can be up to 32 bits long, but can be fewer
	// single rune uses single quotes

	e := 'a'
	fmt.Printf("%v, %T\n", e, e)
	// need to use ReadRune() to convert it back to UTF32 character

	// CONSTANTS
	// can be made up of any of the primitive types
	// Can't be anything mutable i.e. no arrays or whatever
	const myConst int = 42
	const inferredType = 24
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", inferredType, inferredType)

	const (
		aIota = iota
		bIota = iota
		cIota = iota
	)

	fmt.Printf("%v\n", aIota)
	fmt.Printf("%v\n", bIota)
	fmt.Printf("%v\n", cIota)

	const (
		dIota = iota
		eIota
		fIota
	)

	fmt.Printf("%v\n", dIota)
	fmt.Printf("%v\n", eIota)
	fmt.Printf("%v\n", fIota)
}
