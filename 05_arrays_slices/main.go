package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "Mango"
	fruitArr[1] = "Grapes"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	fruitArr2 := [2]string{"Mango", "Grapes"}
	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[0])

	fruitSlice := []string{"Mango", "Grapes", "Peaches"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
