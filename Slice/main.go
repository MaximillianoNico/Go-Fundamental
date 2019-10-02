package main

import "fmt"

func main() {
	// dynamic array
	var fruit = []string{"Apple", "Durian", "Melon"}

	fruit = append(fruit, "Manggo")
	fmt.Println("Total Length Array : ", len(fruit))
	fmt.Println("List Array : ", fruit)
}
