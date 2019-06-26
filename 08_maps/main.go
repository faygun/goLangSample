package main

import "fmt"

func main() {
	emails := map[string]string{"test": "test@test.com", "test2": "test2@test.com", "test3": "test3@test.com"}
	// emails := make(map[string]string)
	// emails["test"] = "test@test.com"
	// emails["test2"] = "test2@test.com"
	// emails["test3"] = "test3@test.com"

	//DELETE
	delete(emails, "test2")

	fmt.Println(emails)
	fmt.Println(emails["test"])
	fmt.Println(len(emails))

}
