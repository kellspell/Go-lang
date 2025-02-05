package main

import (
	"fmt"
)

/*

This is a simple program that demonstrates how to declare variables in Go
There are two ways to declare a variable in Go
One here outside of a function and the other inside a function

*/

var z int = 15

// Here is an example of declaring variables as a block
var (
	actorName    string = "Elizabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {

	// Here is easy way to declare a variable
	var x int = 5

	// Because the compiler can infer the type of the variable
	// you can also declare a variable like this
	y := 10

	fmt.Println(x, y, z)
	fmt.Println(actorName, companion, doctorNumber, season)
}
