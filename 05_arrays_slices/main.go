package main

import "fmt"

func main() {
	//Arrays
	// var fruitSlice [2]string

	// fruitSlice[0] = "Apple"
	// fruitSlice[1] = "Orange"

	fruitSlice := []string{"Orange", "Apple", "Banana", "Cherry", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Printf("Array Count : %d\n", len(fruitSlice))
	fmt.Println(fruitSlice[2:5])
}
