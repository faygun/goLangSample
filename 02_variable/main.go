package main

import "fmt"

var surName string = ""

func main() {
	name, email := "Faruk", "faygun89@gmail.com"
	age := 30
	isMale := true
	surName = "Aygun"

	var size float32 = 2.3

	fmt.Println(name)
	fmt.Println(surName)
	fmt.Println(email)
	fmt.Println(age)
	fmt.Println(isMale)

	fmt.Println(size)
	fmt.Printf("%T\n", age)
}
