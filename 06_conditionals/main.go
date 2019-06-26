package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "green"

	switch color {
	case "red":
		fmt.Println("Colour is red")
	case "blue":
		fmt.Println("Colour is blue")
	default:
		fmt.Println("Colour is not red and blue")
	}

}
