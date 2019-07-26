package main

import "fmt"

func main() {

	// Using `var`
	var age int32 = 23
	const isCool = true

	// Shorthand
	name := "Tanay"

	x, y := "some_x", "some_ y"

	fmt.Println(name, age, isCool, x, y)
	// Print type of a variable
	fmt.Printf("%T\n", name)
}
