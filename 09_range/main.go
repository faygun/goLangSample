package main

import (
	"fmt"
)

func main() {
	// ids := []int{1, 4, 6, 88, 5, 32, 11}

	// for i, id := range ids {
	// 	fmt.Printf("%d ID = %d\n", i, id)
	// }

	// for _, id := range ids {
	// 	fmt.Printf("ID = %d\n", id)
	// }

	// sum := 0
	// for _, id := range ids {
	// 	sum += id
	// }

	// fmt.Println(sum)

	emails := map[string]string{"test": "test@test.com", "test2": "test2@test.com", "test3": "test3@test.com"}

	// for k, v := range emails {
	// 	fmt.Printf("%s - %s\n", k, v)
	// }

	for _, v := range emails {
		fmt.Printf("name : %s\n", v)
	}
}
