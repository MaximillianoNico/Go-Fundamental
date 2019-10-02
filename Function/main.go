package main

import "fmt"

func main() {
	const value = 10
	fmt.Println("Additional : ", displayMessage(value, value))

	// multiple return function
	var x1, x2 = displayMessageMultiple(value, value)
	fmt.Println("Multiple return  : ", x1)
	fmt.Println("Multiple return  : ", x2)
}

func displayMessage(x int, y int) int {
	var z = x + y
	return z
}

func displayMessageMultiple(x int, y int) (int, int) {
	var z = x * y
	var a = x / y

	return z, a
}
