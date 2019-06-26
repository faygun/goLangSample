package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greeting() string {
	return "Hello I am " + p.firstName + " " + p.lastName + " I am " + strconv.Itoa(p.age) + " years old"
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMaried(newSurName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = newSurName
	}
}

func main() {
	person1 := Person{"Faruk", "Aygun", "Balıkesir", "m", 30}

	// fmt.Println(person1)
	person1.hasBirthday()
	person1.getMaried("Kılıc")
	fmt.Print(person1.greeting())
}
