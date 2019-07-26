package main

import "fmt"

func main() {
	a := 5
	b := 10

	if a <= b {
		fmt.Printf("%d is less than or equal to %d\n", a, b)
	} else {
		fmt.Printf("%d is less than %d\n", b, a)
	}

	color := "green"

	if color == "red" {
		fmt.Println("color is red!")
	} else if color == "blue" {
		fmt.Println("color is blue!")
	} else {
		fmt.Println("color is not blue or red!")
	}

	switch color {
	case "red":
		fmt.Println("color is red!")
	case "blue":
		fmt.Println("color is blue!")
	default:
		fmt.Println("color is not blue or red!")
	}
}
