package main

import "fmt"

func main() {
	ids := []int{33, 34, 56, 76, 11}

	// loop throught ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum:", sum)

	// range with map
	fmt.Println("\nRange with Map")
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s:%s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
