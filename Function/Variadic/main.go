package main

import "fmt"

func main() {
	var average = calculateAverage(1, 1, 2, 4, 5, 6, 7)
	var message = fmt.Sprintf("Average : %.2f ", average)

	fmt.Printf(message)
}

func calculateAverage(number ...int) float64 {
	var total int = 0

	for _, n := range number {
		total += n
		fmt.Println(total)
	}

	var avg = float64(total) / float64(len(number))
	return avg
}
