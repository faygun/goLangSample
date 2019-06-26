package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	*b = 10

	fmt.Println(a)
}
