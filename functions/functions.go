package main

import "fmt"

// We are not allowed to create a shadow variable
// shadow variable is a variable that has the same name as a variable in the outer scope
// var a int = 60 -> this will throw an error because we have a variable a in the outer scope
// this is shadowing variable

// Lets talk about iota
// iota is a special constant that is used to generate sequential values
// iota is used to generate sequential values
const (
	e = iota
	f = iota
	g = iota
)

// We also could use iota to generate sequential values
// iota is used to generate sequential values
const (
	a = iota
	b
	c
)

// it will generate the following values
// a = 0
// b = 1
// c = 2

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist

	// we could also use the _ to skip a value the initial value will be the previous value + 1
	// _= iota
)

func main() {
	// In this lesson we will learn about functions in Go
	// Functions are blocks of code that perform a specific task
	// Functions can be used to perform a specific task
	// lets start by creating a constant variable
	/*
		const myConst int = 42
		fmt.Printf("%v, %T\n", myConst, myConst)

		// constant types
		const a int = 14
		const b string = "foo"
		const c float32 = 3.14
		const d bool = true
		fmt.Printf("%v, %T\n", a, a)
		fmt.Printf("%v, %T\n", b, b)
		fmt.Printf("%v, %T\n", c, c)
		fmt.Printf("%v, %T\n", d, d)

		// printing iota
		fmt.Printf("%v,\n", e)
		fmt.Printf("%v,\n", f)
		fmt.Printf("%v,\n", g)

	*/

	// Lets to create a example of how to implement this in a application
	var specialistType int = catSpecialist
	fmt.Printf("%v,\n", specialistType == catSpecialist)

}
