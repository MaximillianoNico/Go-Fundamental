package main

import "fmt"

func main() {
	var price = map[string]int{"chicken_nugget": 2000, "sate": 3000}
	fmt.Println("Sate ayam ", price["sate"])
}
