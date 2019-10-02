package main

import "fmt"

func main() {
	var fruit = [3]string{"Apple", "Durian", "Melon"}
	// replace item array
	fruit[1] = "Manggo"
	fmt.Println("Total Length Array : ", len(fruit))
	fmt.Println("List Array : ", fruit)
}
