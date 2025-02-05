package main

import (
	"fmt"
)

// In this lesson we will learn about the basic types in Go
/*
	1. int
	2. float
	3. string
	4. bool
	5. complex
*/

func main() {
	// Boolean type
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	// comparison
	x := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", m, m)

	// Integer type
	var a int = 10
	var b int8 = 10
	var c int16 = 10
	var d int32 = 10
	var e int64 = 10
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	// math operation
	f := 10
	g := 3
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", f+g, f+g)
	fmt.Printf("%v, %T\n", f-g, f-g)
	fmt.Printf("%v, %T\n", f*g, f*g)
	fmt.Printf("%v, %T\n", f/g, f/g)

	// Float type
	var h float32 = 3.14
	var i float64 = 13.7e72
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)

	// String type
	var j string = "Hello"
	var k string = "World"
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// String concatenation
	var l string = j + " " + k
	fmt.Printf("%v, %T\n", l, l)

	// String interpolation
	var o string = fmt.Sprintf("%v %v", j, k)
	fmt.Printf("%v, %T\n", o, o)

	// Complex type
	var p complex64 = 1 + 2i
	var q complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("%v, %T\n", q, q)
}
