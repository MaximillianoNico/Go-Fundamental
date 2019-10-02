package main

import "fmt"

func main() {
	var a = 10
	var b = 12

	fmt.Println("[additional] Total a + b = %d", a+b)
	fmt.Println("[subtraction] Total a - b = %d", a-b)
	fmt.Println("[division] Total a : b = %d", a/b)
	fmt.Println("[multiplication] Total a x b = %d", a*b)
	fmt.Println("[modulus] Total = %d", a%b)
}
