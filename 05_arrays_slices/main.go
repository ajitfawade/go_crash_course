package main

import (
	"fmt"
)

func main() {
	// Array
	// var fruitArr [2]string
	// assign values

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign at the same time

	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
