package main

import "fmt"

func main() {
	var value = 8

	// condition if - else
	fmt.Println("If Else")
	if value <= 10 && value >= 8 {
		fmt.Println("Grade A")
	} else if value < 8 && value >= 6 {
		fmt.Println("Grade B")
	} else if value < 6 {
		fmt.Println("Grade C")
	}

	// condition switch - case
	fmt.Println("Switch Case")
	switch value {
	case 10:
		fmt.Println("Grade A")
	case 8:
		fmt.Println("Grade B")
	case 6:
		fmt.Println("Grade C")
	}
}
